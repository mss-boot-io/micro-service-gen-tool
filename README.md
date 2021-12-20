# micro-service-gen-tool
White Matrix Micro Service Generate CLI Tool

## usage
### linux
```bazaar
## linux
curl -O https://whitematrixtech.github.io/micro-service-gen-tool/latest/linux_amd64.tar.gz
## mac
curl -O https://whitematrixtech.github.io/micro-service-gen-tool/latest/darwin_amd64.tar.gz
##
tar -zxvf linux_amd64.tar.gz
## create config local
cat >> config.yml <<eof
service: proto-demo
templateUrl: https://github.com/lwnmengjing/tempate-demo
createRepo: false
destination: ./
ignore:
  - "static.txt"
data:
  service: proto-demo
eof
## create config for github
cat >> config.yml <<eof
service: proto-demo
templateUrl: https://github.com/lwnmengjing/tempate-demo
createRepo: true
destination: ./
ignore:
  - "static.txt"
data:
  service: proto-demo
github:
  token: {github_token}
  description: description
  organization: WhiteMatrixTech
eof
# generate code
./generate-tool --config=config.yml
```