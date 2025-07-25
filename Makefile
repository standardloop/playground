include Makefile.properties

all: check cluster infra app

setup: cluster infra

# kubectl config unset current-context
# clean:  app.clean infra.clean cluster.clean

clean: cluster.clean

check: check.dependencies check.docker

check.docker:
	@docker info > /dev/null 2>&1

check.dependencies:
	@for dep in $(DEPENDENCIES); do \
    	if ! which $$dep > /dev/null; then\
   			echo $$dep not found;\
			exit 1;\
		fi\
    done

cluster:
	kind create cluster --config=./kind-config.yaml

cluster.clean:
	kind delete cluster --name $(CLUSTER_NAME)

cluster.context:
	kubectx $(CLUSTER_CONTEXT)

cluster.context.info:
	kubectl cluster-info --context $(CLUSTER_CONTEXT)

infra: infra.ingress infra.prometheus infra.metrics infra.istio #infra.argocd 

infra.clean: infra.prometheus.clean infra.ingress.clean

infra.upgrade: infra.ingress.upgrade infra.prometheus.upgrade

infra.redis:
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm repo update
	kubectl create namespace redis
	helm install my-redis bitnami/redis -n redis

infra.ingress:
	-helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
	-helm repo update
	-kubectl create namespace $(INGRESS_NAMESPACE)
	helm install ingress-nginx ingress-nginx/ingress-nginx -n $(INGRESS_NAMESPACE) --values ./deploy/helm/ingress-nginx.yaml --version 4.0.6
	kubectl wait --namespace $(INGRESS_NAMESPACE) \
		--for=condition=ready pod \
		--selector=app.kubernetes.io/component=controller \
		--timeout=180s

infra.ingress.upgrade:
	helm upgrade ingress-nginx ingress-nginx/ingress-nginx -n $(INGRESS_NAMESPACE) --values ./deploy/helm/ingress-nginx.yaml --version 4.0.6

infra.ingress.clean:
	helm uninstall ingress-nginx -n $(INGRESS_NAMESPACE)

infra.istio: infra.istio.base infra.istio.gateway infra.istio.resources

infra.istio.base:
	-helm repo add istio https://istio-release.storage.googleapis.com/charts
	-helm repo update
	-kubectl create namespace istio-system
	helm install istio-base istio/base -n istio-system
	helm install istiod istio/istiod -n istio-system --wait

infra.istio.gateway:
	-kubectl create namespace istio-gateway
	helm install istio-gateway istio/gateway -n istio-gateway --values ./deploy/helm/istio-gateway.yaml

infra.istio.resources:
	-kubectl apply -k deploy/kustomize/istio

infra.prometheus:
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo update
	kubectl create namespace kube-prometheus-stack
	helm install kube-prometheus-stack prometheus-community/kube-prometheus-stack -n $(PROM_STACK_NAMESPACE) --values ./deploy/helm/kube-prometheus-stack.yaml --version $(PROM_STACK_VERSION)

infra.prometheus.upgrade:
	helm upgrade kube-prometheus-stack prometheus-community/kube-prometheus-stack -n $(PROM_STACK_NAMESPACE) --values ./deploy/helm/kube-prometheus-stack.yaml --version $(PROM_STACK_VERSION)

infra.prometheus.password:
	@echo "prom-operator"

infra.prometheus.clean:
	helm uninstall kube-prometheus-stack 
	kubectl delete crd alertmanagerconfigs.monitoring.coreos.com -n $(PROM_STACK_NAMESPACE)
	kubectl delete crd alertmanagers.monitoring.coreos.com -n $(PROM_STACK_NAMESPACE)
	kubectl delete crd podmonitors.monitoring.coreos.com -n $(PROM_STACK_NAMESPACE)
	kubectl delete crd probes.monitoring.coreos.com -n $(PROM_STACK_NAMESPACE)
	kubectl delete crd prometheuses.monitoring.coreos.com -n $(PROM_STACK_NAMESPACE)
	kubectl delete crd prometheusrules.monitoring.coreos.com -n $(PROM_STACK_NAMESPACE)
	kubectl delete crd servicemonitors.monitoring.coreos.com -n $(PROM_STACK_NAMESPACE)
	kubectl delete crd thanosrulers.monitoring.coreos.com -n $(PROM_STACK_NAMESPACE)

infra.argocd:
	kubectl apply -k deploy/kustomize/apps/argocd/dev

infra.argocd.password:
	@kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo

infra.argocd.clean:
	kubectl delete -k deploy/kustomize/apps/argocd/dev

infra.metrics:
	helm repo add metrics-server https://kubernetes-sigs.github.io/metrics-server/
	helm upgrade --install metrics-server metrics-server/metrics-server -n kube-system --values ./deploy/helm/metric-server.yaml

app: api db

db: db.deploy

db.deploy:
	kubectl apply -k deploy/kustomize/apps/mysql/dev
	kubectl apply -k deploy/kustomize/apps/postgres/dev
	kubectl apply -k deploy/kustomize/apps/mongo/dev

api: api.build api.deploy

api.build:
	docker build -t api:latest -t api:$(API_TAG) src/api
	kind load docker-image api:$(API_TAG) --name $(CLUSTER_NAME)

api.deploy:
	cd deploy/kustomize/apps/api/dev && kustomize edit set image api:$(API_TAG) 
	kubectl apply -k deploy/kustomize/apps/api/dev

api.uninstall:
	kubectl delete -k deploy/kustomize/apps/api/dev

api.test.basic:
	curl http://api.local:80/api/v1/health | jq

api.test.rand:
	curl http://api.local:80/api/v1/rand | jq

docker.clean:
	docker stop $(docker ps -a -q)
	docker rm $(docker ps -a -q)

docker.compose:
	docker compose build
	docker compose up -d
