# global
ci_settings(timeout='60m', readiness_timeout='10m')
default_registry('localhost:5001')
allow_k8s_contexts('kind-slke-1')

# app
docker_build('localhost:5001/api', 'src/api', dockerfile='src/api/Dockerfile')
k8s_yaml(kustomize('deploy/apps/'))
