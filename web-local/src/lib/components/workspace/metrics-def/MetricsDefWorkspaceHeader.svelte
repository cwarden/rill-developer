<script lang="ts">
  import { goto } from "$app/navigation";
  import MetricsIcon from "@rilldata/web-common/components/icons/Metrics.svelte";
  import { notifications } from "@rilldata/web-common/components/notifications";
  import { useRuntimeServiceRenameFileAndReconcile } from "@rilldata/web-common/runtime-client";
  import { runtimeStore } from "@rilldata/web-local/lib/application-state-stores/application-store";
  import { useQueryClient } from "@sveltestack/svelte-query";
  import {
    isDuplicateName,
    renameFileArtifact,
    useAllNames,
  } from "../../../svelte-query/actions";
  import { EntityType } from "../../../temp/entity";
  import MetricsDefinitionExploreMetricsButton from "../../metrics-definition/MetricsDefinitionExploreMetricsButton.svelte";
  import WorkspaceHeader from "../core/WorkspaceHeader.svelte";

  export let metricsDefName;
  export let metricsInternalRep;

  $: runtimeInstanceId = $runtimeStore.instanceId;
  $: allNamesQuery = useAllNames(runtimeInstanceId);
  const queryClient = useQueryClient();
  const renameMetricsDef = useRuntimeServiceRenameFileAndReconcile();

  const onChangeCallback = async (e) => {
    if (!e.target.value.match(/^[a-zA-Z_][a-zA-Z0-9_]*$/)) {
      notifications.send({
        message:
          "Dashboard name must start with a letter or underscore and contain only letters, numbers, and underscores",
      });
      e.target.value = metricsDefName; // resets the input
      return;
    }
    if (isDuplicateName(e.target.value, $allNamesQuery.data)) {
      notifications.send({
        message: `Name ${e.target.value} is already in use`,
      });
      e.target.value = metricsDefName; // resets the input
      return;
    }

    try {
      const toName = e.target.value;
      const type = EntityType.MetricsDefinition;
      await renameFileArtifact(
        queryClient,
        runtimeInstanceId,
        metricsDefName,
        toName,
        type,
        $renameMetricsDef
      );
      goto(`/dashboard/${toName}/edit`, { replaceState: true });
    } catch (err) {
      console.error(err.response.data.message);
    }
  };

  $: titleInput = metricsDefName;
</script>

<WorkspaceHeader
  {...{ titleInput, onChangeCallback }}
  showStatus={false}
  showInspectorToggle={false}
>
  <MetricsIcon slot="icon" />
  <MetricsDefinitionExploreMetricsButton
    slot="cta"
    {metricsDefName}
    {metricsInternalRep}
  />
</WorkspaceHeader>
