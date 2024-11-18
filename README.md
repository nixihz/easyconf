# easyconf

Allow to generate config file from yaml template and data.

### Installation

```shell
go install github.com/nixihz/easyconf

```
### Usage

params  
- -r local
  - Type of generated file: local, k8s  
- -ta common
  - Format of yaml file: common, configmap
- -t ./example/configmap_tpl.yaml
  - template file 
- -f ./example/config-dev.yaml
  -  date file 
- -o ./example/
  - The directory or file where the file is generated.
    - -r local, -o is directory
    - -r k8s，-o is a file path

### Example

```shell
# convert to local yaml file
easyconf -r local -ta common -t ./configmap_tpl.yaml -f ./config-dev.yaml -o ./configs/

# generate k8s configmap
easyconf -r k8s -ta configmap -t ./configmap_tpl.yaml -f ./config-dev.yaml -o ./configs/configmap.yaml

```
