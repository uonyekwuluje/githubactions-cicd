# GitHub Actions Recipes
This is a collection of configs on setting up GitHub Actions and Pipelines.

### **Runner Token Generation**
Follow the steps below to generate your runner access token
* Loginto your GitHub Account
* Click on your repository and click on settings
* Click on Actions
* Click on Actions Permissions. In this cas ```Allow all actions```
* Click on ```Add runner```
* ***NOTE:*** *Note the string in configure*


### **Repository Preparation**
* Create the folder ```.github/workflows``` in your repository
* Create your **yaml** file in this folder. This is your pipeline 

### **Self Hostaed Runner CentOS 7**
To install Self Hosted Runners, visit [Self Hosted Runner](https://github.com/actions/runner)
```
sudo yum update -y
sudo yum group install -y 'Development Tools'
sudo yum install -y libxml2 libxml2-devel libxslt libxslt-devel \
libffi-devel openssl-devel make openssl-devel bzip2-devel gcc wget
sudo yum install -y java-11-openjdk-devel java-11-openjdk zlib \
lttng-ust openssl-libs krb5-libs libicu

GITHUBACTIONS_RUNNER="2.275.1"
SYSTEM_USER="centos"
sudo mkdir /opt/actions-runner 
sudo chown ${SYSTEM_USER}:${SYSTEM_USER} -R /opt/actions-runner/
cd /opt/actions-runner
curl -O -L https://github.com/actions/runner/releases/download/v${GITHUBACTIONS_RUNNER}/actions-runner-linux-x64-${GITHUBACTIONS_RUNNER}.tar.gz
tar xzf ./actions-runner-linux-x64-${GITHUBACTIONS_RUNNER}.tar.gz
rm actions-runner-linux-x64-${GITHUBACTIONS_RUNNER}.tar.gz
```

### **Self Hostaed Runner Ubintu 12.04 LTS**
To install Self Hosted Runners, visit [Self Hosted Runner](https://github.com/actions/runner)
```
sudo apt update && sudo apt upgrade -y
sudo apt install -y build-essential
sudo apt install -y autoconf automake gdb git libffi-dev zlib1g-dev libssl-dev

GITHUBACTIONS_RUNNER="2.275.1"
SYSTEM_USER="ubuntu"
sudo mkdir /opt/actions-runner
sudo chown ${SYSTEM_USER}:${SYSTEM_USER} -R /opt/actions-runner/
cd /opt/actions-runner
curl -O -L https://github.com/actions/runner/releases/download/v${GITHUBACTIONS_RUNNER}/actions-runner-linux-x64-${GITHUBACTIONS_RUNNER}.tar.gz
tar xzf ./actions-runner-linux-x64-${GITHUBACTIONS_RUNNER}.tar.gz
rm actions-runner-linux-x64-${GITHUBACTIONS_RUNNER}.tar.gz
```

For install instructions, check Actions setup for the given repository. After setup, set as service
```
cd /opt/actions-runner
./config.sh --url https://github.com/<account>/<repository> --token <token>
```
The sequence of steps will be
```
--------------------------------------------------------------------------------
|        ____ _ _   _   _       _          _        _   _                      |
|       / ___(_) |_| | | |_   _| |__      / \   ___| |_(_) ___  _ __  ___      |
|      | |  _| | __| |_| | | | | '_ \    / _ \ / __| __| |/ _ \| '_ \/ __|     |
|      | |_| | | |_|  _  | |_| | |_) |  / ___ \ (__| |_| | (_) | | | \__ \     |
|       \____|_|\__|_| |_|\__,_|_.__/  /_/   \_\___|\__|_|\___/|_| |_|___/     |
|                                                                              |
|                       Self-hosted runner registration                        |
|                                                                              |
--------------------------------------------------------------------------------

# Authentication


√ Connected to GitHub

# Runner Registration

Enter the name of runner: [press Enter for githubactions-agent1] 

This runner will have the following labels: 'self-hosted', 'Linux', 'X64' 
Enter any additional labels (ex. label-1,label-2): [press Enter to skip] java,python

√ Runner successfully added
√ Runner connection is good

# Runner settings

Enter name of work folder: [press Enter for _work] 

√ Settings Saved.
```

Run the command below to test
```
./run.sh
```
Check the Actions menu and you should see the runner on-line. After this, kill this commanmd and set it up as a service
```
sudo ./svc.sh install
sudo ./svc.sh start
sudo ./svc.sh status
```

### **Repository Release**
* [Github Action Releases](https://github.com/actions/runner/releases)
