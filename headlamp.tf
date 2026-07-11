# /*
#     WIP
#     - kubectl -n kube-system create serviceaccount headlamp-admin  # wip
#     - kubectl create clusterrolebinding headlamp-playground-admin --serviceaccount=kube-system:headlamp-admin --clusterrole=cluster-admin
#     - kubectl create token headlamp-admin -n kube-system | pbcopy
# */

# data "kustomization_build" "headlamp" {
#   path = "deploy/infrastructure/headlamp"
# }

# resource "kustomization_resource" "headlamp" {
#   for_each   = data.kustomization_build.headlamp.ids
#   manifest   = data.kustomization_build.headlamp.manifests[each.value]
#   depends_on = [helm_release.flux2, kustomization_resource.flux_sources]
# }

# resource "null_resource" "wait_for_headlamp" {
#   depends_on = [kustomization_resource.headlamp]

#   provisioner "local-exec" {
#     command = "kubectl wait helmrelease/headlamp --for=condition=Ready --timeout=${var.HELM_RELEASE_TIMEOUT} -n kube-system"
#   }
# }
