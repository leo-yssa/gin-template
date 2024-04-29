version=1.21

sudo apt remove --autoremove golang-go; 
sudo apt install snapd; 
sudo snap install --classic --channel=${version}/stable go