## Preparation

```bash
kind create cluster --name 1 --config <(cat << EOF
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
  - role: control-plane
    extraPortMappings:
    - containerPort: 30599
      hostPort: 8000
  - role: worker
EOF
)
```

deploy Nginx 

```bash
kubectl create deployment nginx --image=nginx

cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Service
metadata:
  labels:
    app: nginx
  name: nginx-svc
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 80
    nodePort: 30599
  selector:
    app: nginx
  type: NodePort
EOF
```
