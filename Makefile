include Makefile.properties

all: check cluster infra app

setup: cluster infra

#clean:  app.clean infra.clean cluster.clean
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
	kind create cluster --config=./kind-config.yaml --name $(CLUSTER_NAME)

cluster.clean:
	kind delete cluster --name $(CLUSTER_NAME)

cluster.context:
	kubectx $(CLUSTER_CONTEXT)

cluster.context.info:
	kubectl cluster-info --context $(CLUSTER_CONTEXT)


infra: infra.ingress infra.prometheus infra.metrics #infra.argocd 

infra.clean: infra.prometheus.clean infra.ingress.clean

infra.upgrade: infra.ingress.upgrade infra.prometheus.upgrade

infra.ingress:
	helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
	helm repo update
	-kubectl create namespace $(INGRESS_NAMESPACE)
	helm install ingress-nginx ingress-nginx/ingress-nginx -n $(INGRESS_NAMESPACE) --values ./deploy/ingress-nginx-config.yaml --version 4.0.6
	kubectl wait --namespace $(INGRESS_NAMESPACE) \
		--for=condition=ready pod \
		--selector=app.kubernetes.io/component=controller \
		--timeout=180s

infra.ingress.upgrade:
	helm upgrade ingress-nginx ingress-nginx/ingress-nginx -n $(INGRESS_NAMESPACE) --values ./deploy/ingress-nginx-config.yaml --version 4.0.6

infra.ingress.clean:
	helm uninstall ingress-nginx -n $(INGRESS_NAMESPACE)

infra.prometheus:
	helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
	helm repo update
	kubectl create namespace kube-prometheus-stack
	helm install kube-prometheus-stack prometheus-community/kube-prometheus-stack -n kube-prometheus-stack --values ./deploy/kube-prometheus-stack-config.yaml --version 36.6.0

infra.prometheus.upgrade:
	helm upgrade kube-prometheus-stack prometheus-community/kube-prometheus-stack -n kube-prometheus-stack --values ./deploy/kube-prometheus-stack-config.yaml --version 36.6.0

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
	kubectl apply -k deploy/argocd/dev

infra.argocd.password:
	@kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d; echo

infra.argocd.clean:
	kubectl delete -k deploy/argocd/dev

infra.metrics:
	helm repo add metrics-server https://kubernetes-sigs.github.io/metrics-server/
	helm upgrade --install metrics-server metrics-server/metrics-server -n kube-system --values ./deploy/metric-server-config.yaml

infra.fleet:
	helm -n fleet-system install --create-namespace --wait fleet-crd https://github.com/rancher/fleet/releases/download/v$(FLEET_VERSION)/fleet-crd-$(FLEET_VERSION).tgz
	helm -n fleet-system install --create-namespace --wait fleet https://github.com/rancher/fleet/releases/download/v$(FLEET_VERSION)/fleet-$(FLEET_VERSION).tgz


app: api ui db

db: db.deploy

db.deploy:
	kubectl apply -k deploy/mysql/dev
	kubectl apply -k deploy/postgres/dev

ui: ui.build ui.deploy

ui.build:
	docker build -t ui:latest -t ui:$(UI_TAG) src/ui
	kind load docker-image ui:$(UI_TAG) --name $(CLUSTER_NAME)

ui.deploy:
	cd deploy/ui/dev && kustomize edit set image ui:$(UI_TAG) 
	kubectl apply -k deploy/ui/dev

ui.dockerrun:
	docker run -p 3000:3000 ui:0.0.8 \
		-e API_PROTOCOOL='http' \
		-e API_EXTERNAL_URL='localhost' \
		-e API_INTERNAL_URL='localhost' \
		-e API_PORT='8080'


ui.uninstall:
	kubectl delete -k deploy/ui/dev

api: api.build api.deploy

api.build:
	docker build -t api:latest -t api:$(API_TAG) src/api
	kind load docker-image api:$(API_TAG) --name $(CLUSTER_NAME)

api.deploy:
	cd deploy/api/dev && kustomize edit set image api:$(API_TAG) 
	kubectl apply -k deploy/api/dev

api.uninstall:
	kubectl delete -k deploy/api/dev

api.test.basic:
	curl http://api.local:80/api/v1/health | jq

api.test.rand:
	curl http://api.local:80/api/v1/rand | jq

docker.clean:
	docker stop $(docker ps -a -q)
	docker rm $(docker ps -a -q)
