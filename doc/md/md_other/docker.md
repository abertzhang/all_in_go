# docker安装

## Mac

[官网](https://www.docker.com/)

## 阿里镜像加速

1. 登录阿里云官网

2. 容器镜像服务

3. 镜像工具/镜像加速器

   ```
   {
     "registry-mirrors": ["https://coda81p9.mirror.aliyuncs.com"]
   }
   ```

## kubernetes

### 下载阿里k8s

[github阿里k8s下载](https://github.com/denverdino/k8s-for-docker-desktop)

### docker启用k8s

### k8s控制台

```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.5.1/aio/deploy/recommended.yaml
//在阿里下载k8s的目录下
kubectl apply -f kubernetes-dashboard.yaml
//检查 kubernetes-dashboard 应用状态
kubectl get pod -n kubernetes-dashboard
//
kubectl proxy
//访问如下地址
http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/
```

[Kubernetes dashboard](http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/)

# 常用命令

## docker

```
docker version
docker images

```

## k8s

```
kubectl cluster-info
//
kubectl get nodes
//
kubectl get pods -n kube-system
```



# 参考文档

[如何安装Docker Desktop for Mac](https://www.bilibili.com/video/BV16E411u75h/?spm_id_from=333.337.search-card.all.click&vd_source=b9aff273129955972ba5e761af57d33d)
