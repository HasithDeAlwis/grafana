import { cloneDeep } from 'lodash';

import { t } from '@grafana/i18n';
import { locationService } from '@grafana/runtime';
import { Alert, LoadingPlaceholder } from '@grafana/ui';
import { RuleIdentifier, RuleWithLocation } from 'app/types/unified-alerting';
import { RulerRuleDTO } from 'app/types/unified-alerting-dto';

import { AlertRuleForm } from '../components/rule-editor/alert-rule-form/AlertRuleForm';
import { useRuleWithLocation } from '../hooks/useCombinedRule';
import { generateCopiedName } from '../utils/duplicate';
import { stringifyErrorLike } from '../utils/misc';
import { rulerRuleToFormValues } from '../utils/rule-form';
import { getRuleName, rulerRuleType } from '../utils/rules';
import { createRelativeUrl } from '../utils/url';

export function CloneRuleEditor({ sourceRuleId }: { sourceRuleId: RuleIdentifier }) {
  const { loading, result: rule, error } = useRuleWithLocation({ ruleIdentifier: sourceRuleId });

  if (loading) {
    return <LoadingPlaceholder text={t('alerting.clone-rule-editor.text-loading-the-rule', 'Loading the rule...')} />;
  }

  if (rule) {
    const ruleClone = cloneRuleDefinition(rule);
    const formPrefill = rulerRuleToFormValues(ruleClone);

    return <AlertRuleForm prefill={formPrefill} />;
  }

  if (error) {
    return (
      <Alert title={t('alerting.clone-rule-editor.title-error', 'Error')} severity="error">
        {stringifyErrorLike(error)}
      </Alert>
    );
  }

  return (
    <Alert
      title={t('alerting.clone-rule-editor.title-cannot-exist', 'Cannot copy the rule. The rule does not exist')}
      buttonContent="Go back to alert list"
      onRemove={() => locationService.replace(createRelativeUrl('/alerting/list'))}
    />
  );
}

function changeRuleName(rule: RulerRuleDTO, newName: string) {
  if (rulerRuleType.grafana.rule(rule)) {
    rule.grafana_alert.title = newName;
  }
  if (rulerRuleType.dataSource.alertingRule(rule)) {
    rule.alert = newName;
  }

  if (rulerRuleType.dataSource.recordingRule(rule)) {
    rule.record = newName;
  }
}

export function cloneRuleDefinition(rule: RuleWithLocation<RulerRuleDTO>) {
  const ruleClone = cloneDeep(rule);
  changeRuleName(
    ruleClone.rule,
    generateCopiedName(getRuleName(ruleClone.rule), ruleClone.group.rules.map(getRuleName))
  );

  if (rulerRuleType.grafana.rule(ruleClone.rule)) {
    ruleClone.rule.grafana_alert.uid = '';

    // Provisioned alert rules have provisioned alert group which cannot be used in UI
    if (Boolean(ruleClone.rule.grafana_alert.provenance)) {
      ruleClone.group = { name: '', rules: ruleClone.group.rules };
    }
  }

  return ruleClone;
}
