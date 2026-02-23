# global
ci_settings(timeout='60m', readiness_timeout='10m')

# app
docker_build('api', 'src/api', dockerfile='src/api/Dockerfile')
k8s_yaml(kustomize('deploy/apps/'))
