
#-*- mode: Python -*-
docker_build('homeapp-go-image', '.', dockerfile='Dockerfile')

k8s_yaml('deployments/postgres.yaml')

k8s_yaml('deployments/kubernetes.yaml')
k8s_resource('homeapp-go', port_forwards=8000, resource_deps=['postgres-statefulset'])

k8s_resource('postgres-statefulset', port_forwards=5432)


