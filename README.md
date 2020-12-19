# GitHub Actions Recipes
This is a collection of configs on setting up GitHub Actions and Pipelines

### **Repository Preparation**
* Create the folder ```.github/workflows``` in your repository
* Create your **yaml** file in this folder. This is your pipeline 

### **Self Hostaed Runner**
To install Self Hosted Runners, visit [Self Hosted Runner](https://github.com/actions/runner)
I'll be using CentOS 7 for my Self Hosted Runner:
```
sudo yum update -y
sudo yum group install -y 'Development Tools'
sudo yum install -y libxml2 libxml2-devel libxslt libxslt-devel \
libffi-devel openssl-devel make openssl-devel bzip2-devel gcc wget
sudo yum install -y java-11-openjdk-devel java-11-openjdk zlib \
lttng-ust openssl-libs krb5-libs libicu

# [https://github.com/actions/runner/releases]
GITHUBACTIONS_RUNNER="2.275.1"
SYSTEM_USER="centos"
sudo mkdir /opt/actions-runner && cd /opt/actions-runner
sudo chown centos:centos -R /opt/actions-runner/
sudo curl -O -L https://github.com/actions/runner/releases/download/v${GITHUBACTIONS_RUNNER}/actions-runner-linux-x64-${GITHUBACTIONS_RUNNER}.tar.gz
sudo tar xzf ./actions-runner-linux-x64-${GITHUBACTIONS_RUNNER}.tar.gz
```
For install instructions, check Actions setup for the given repository. After setup, set as service
```
cd /opt/actions-runner
sudo ./svc.sh install
sudo ./svc.sh start
sudo ./svc.sh status
```
