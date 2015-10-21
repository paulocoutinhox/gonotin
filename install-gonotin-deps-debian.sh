echo "> Updating packages..."
apt-get update
echo ""

echo "> Installing C++..."
apt-get install build-essential
apt-get install g++-4.8
update-alternatives --install /usr/bin/g++ g++ /usr/bin/g++-4.8 50
echo ""

echo "> Installing Go..."
apt-get install curl git mercurial make binutils bison gcc
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
source ~/.bashrc
apt-get install bison
gvm install go1.4.3
gvm use go1.4.3
gvm install go1.5.1
gvm use go1.5.1
echo ""

echo "> Installing Python..."
sudo apt-get install python
echo ""
