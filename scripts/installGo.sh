version=1.22.2
base=/usr/local
root=${base}/go
file=go${version}.linux-amd64.tar.gz
if [ ! -f ${file} ] || [ ! -d ${base} ]; then
   echo "File ${file} does not exist."
   wget https://golang.org/dl/${file}
   sudo tar -C ${base} -xzf ${file}
fi

if ! grep  'GOROOT' ~/.bashrc ; then
    echo "export GOROOT=${root}" >> ~/.bashrc
    echo "export PATH=${root}/bin:\$PATH" >> ~/.bashrc
    source ~/.bashrc
fi