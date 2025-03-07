import { goto } from "$app/navigation";
import { notifications } from "@rilldata/web-common/components/notifications";
import type {
  V1PutFileAndReconcileResponse,
  V1ReconcileError,
} from "@rilldata/web-common/runtime-client";
import { fileArtifactsStore } from "@rilldata/web-local/lib/application-state-stores/file-artifacts-store";
import { invalidateAfterReconcile } from "@rilldata/web-local/lib/svelte-query/invalidation";
import { EntityType } from "@rilldata/web-local/lib/temp/entity";
import { getFilePathFromNameAndType } from "@rilldata/web-local/lib/util/entity-mappers";
import type { QueryClient, UseMutationResult } from "@sveltestack/svelte-query";

export async function createSource(
  queryClient: QueryClient,
  instanceId: string,
  tableName: string,
  yaml: string,
  createSourceMutation: UseMutationResult<V1PutFileAndReconcileResponse>
): Promise<V1ReconcileError[]> {
  const resp = await createSourceMutation.mutateAsync({
    data: {
      instanceId,
      path: getFilePathFromNameAndType(tableName, EntityType.Table),
      blob: yaml,
      create: true,
      // create source is used to upload and replace.
      // so we cannot send createOnly=true until we refactor it to use refresh source
      createOnly: false,
      strict: true,
    },
  });

  if (resp.errors.length) {
    return resp.errors;
  }
  goto(`/source/${tableName}`);
  invalidateAfterReconcile(queryClient, instanceId, resp);
  fileArtifactsStore.setErrors(resp.affectedPaths, resp.errors);
  notifications.send({ message: `Created source ${tableName}` });
  return [];
}
