#!/bin/bash

# 根据不同的模板文件 和 公共的资源文件 ，批量替换变量
# 研发改具体参数 只改模板文件，资源文件由运维负责更新和修改



if [ $# -eq 2 ]; then
    tplF="$1"
    resF="$2"
else
    echo "Usage: $0 templateFile<./conf/config.json.temp> resourceFile<./conf/resource.conf>"
    exit 1
fi

[ ! -s $tplF ] && echo "[ERROR] none exists template file $tplF" && exit 1
[ ! -s $resF ] && echo "[ERROR] none exists resource file $resF" && exit 1

CWD=$(cd `dirname $0`; pwd)
cd $CWD

#destName=`echo $tplF | sed 's/.*template//'`
#echo $destName

#### 
tmpD="$CWD/tmp"
[ ! -d $tmpD ] && mkdir -p $tmpD

tmpF="$tmpD/temp_.$RANDOM"
/bin/cp -af $tplF $tmpF >/dev/null
if [ $? -ne 0 ]; then
    exit 1
fi

source $resF

####
for i in `grep -oP "<.+?>" $tplF`
do
    #echo "$i"
    tplVar=${i##<}
    tplVar=${tplVar%>}

    haveVar=$(eval echo \$$tplVar) 
    if [ "x$haveVar" != "x" ] || [ `grep -wc $tplVar $resF` -ge 1 ]; then
        :
        #echo "$i => $tplVar => haveVar:$haveVar"
    else
        echo "[ERROR] none find var $i => $tplVar => haveVar:$haveVar"
        /bin/rm -f $tmpF
        exit 1
    fi
    sed -i "s/$i/$haveVar/" $tmpF 
done

cat $tmpF
#content=$(cat $tmpF)
#echo "$content"
/bin/rm -f $tmpF

