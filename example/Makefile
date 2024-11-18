.PHONY: easyconf-local
# local env using dev config
easyconf-local:
	easyconf -r local -ta common -t ./configmap_tpl.yaml -f ./config-dev.yaml -o ./configs/

.PHONY: easyconf-k8s
# k8s configmap
easyconf-k8s:
	easyconf -r k8s -ta configmap -t ./configmap_tpl.yaml -f ./config-dev.yaml -o ./configs/configmap.yaml

