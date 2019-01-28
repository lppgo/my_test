#!/usr/bin/python
# -*- coding: utf-8 -*-
"""Fabric file to update firmware-upgrade."""
from fabric.api import run, env, local, cd, put
from fabric.contrib import files, console
from datetime import datetime
import time

env.hosts = [
    # 'prd-6.yeeuu.com',
    'test-1.yeeuu.com',
]
# env.key_filename = "~/.ssh/yeeuu_rsa"
env.user = 'root'
now = datetime.now().strftime("%Y%m%d%H%M%S")
built = False


def build(env='darwin'):
    """Local binary."""
    global built
    if not built:
        local('GOOS={0} go build'.format(env))
        built = True


def test():
    """Test this project."""
    local('go test -v ./...')


def update():
    """Upgrade firmware-upgrade capture."""
    build('linux')
    if files.exists('/mnt/code/alms-v3/alms-v3'):
        run('mv /mnt/code/alms-v3/alms-v3 ' +
            '/mnt/code/alms-v3/alms-v3_' + now)
    with cd('/mnt/code/alms-v3'):
        put('alms-v3', 'alms-v3')
        run('chmod +x alms-v3')
    run('supervisorctl restart alms-v3')
    time.sleep(3)
    run('supervisorctl status alms-v3')


from fabric.colors import red, green
import sys

def rollback():
    """ 回滚版本部署 """
    print(green("回滚 alms-v3 项目版本 脚本 慎用 慎用 慎用 !!!!"))
    with cd('/mnt/code/alms-v3'):
        print(red("Note: 最近发布的10个版本---按更改日期排序"))
        run("ls -t | head -10 | grep alms-v3_'2018[0-9]'")
        print(red("请选择要回滚的版本:例如(alms-v3_20180427191110)"))
        version = raw_input(">>>")
        if version == "":
            sys.exit(0)
        if files.exists('/mnt/code/alms-v3/'+version.strip()):
            print(green("版本正确!!!"))
            print(green("当前回滚的版本是: " + version))
            ok = console.confirm(red("是否要执行回滚操作?谨慎操作！！！"), default=False)
            if ok :
                print(green("copy 当前运行的版本!!!"))
                run('cp alms-v3  alms-v3_' + now)
                print(green("正在执行版本回滚命令!!!"))
                run('mv /mnt/code/alms-v3/%s  alms-v3'% version)
                run('supervisorctl restart alms-v3')
                time.sleep(5)
                run('supervisorctl status alms-v3')
            else:
                print(green("退出回滚操作！Bye!"))
                sys.exit(0)
        else:
            print(red("版本不存在, 回滚失败！！！"))

    