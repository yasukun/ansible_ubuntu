# ansible_ubuntu

## Dependencies

 * [ansible-r](https://github.com/Oefenweb/ansible-r)

## usage

/etc/ansible/hosts

```bash
localhost	ansible_connection=local	ansible_ssh_user=hoge
```
command

```bash
$ ansible-galaxy install -p ./roles -r roles.yml
$ ansible-playbook site.yml --ask-vault-pass
```
## set up nginx

```bash
$ cd /etc/nginx/cert
$ sudo openssl genrsa -des3 -out server.key 1024
$ sudo openssl req -new -key server.key -out server.csr
$ sudo cp server.key server.key.org
$ sudo openssl rsa -in server.key.org -out ssl.key
$ sudo openssl x509 -req -days 1825 -in server.csr -signkey ssl.key -out ssl.crt
$ cd /etc/nginx
$ sudo mv nginx.conf nginx.conf.bk
$ sudo mv nginx.conf.main nginx.conf
$ sudo service nginx restart
```

## set up jupyter

```bash
$ pyenv install anaconda3-4.4.0
$ pyenv global anaconda3-4.4.0
```
