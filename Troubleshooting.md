# Production-Style Troubleshooting Scenarios

This document highlights the most valuable troubleshooting scenarios encountered while building and deploying this project. These issues reflect real-world DevOps challenges involving Docker, Kubernetes, AWS, CI/CD, and GitOps.

---

# 1. Docker Image Architecture Mismatch

## Problem

Pods continuously failed to start after deployment to Amazon EKS.

```text
exec ./main: exec format error
```

## Root Cause

The Docker image was built on an Apple Silicon (ARM64) MacBook, while the EKS worker nodes were running on AMD64 EC2 instances.

The binary architecture inside the image was incompatible with the worker nodes.

## Troubleshooting Steps

* Checked pod status.
* Inspected pod logs.
* Identified `exec format error`.
* Verified Docker image architecture.

## Resolution

Configured Docker Buildx to build a multi-platform image.

```bash
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  --push \
  -t <ECR_REPOSITORY>:<TAG> .
```

## Learning

Always build multi-architecture images when developing on Apple Silicon and deploying to AMD64 environments.

---

# 2. GitHub Actions Built Successfully but Images Were Not Updated

## Problem

The CI workflow completed successfully, but Amazon ECR continued to show the previous image version.

## Root Cause

The workflow pushed Docker images only from the `main` branch.

Testing was being performed on a feature branch.

## Troubleshooting Steps

* Verified workflow trigger configuration.
* Checked GitHub Actions logs.
* Compared image tags in Amazon ECR.

## Resolution

Merged the Pull Request into `main`, allowing the production workflow to publish the new image.

## Learning

Always understand the trigger conditions configured in GitHub Actions.

---

# 3. Kubernetes Could Not Pull Images from Amazon ECR

## Problem

Application pods remained in `ImagePullBackOff`.

## Root Cause

The image either did not exist in Amazon ECR or Kubernetes did not have permission to pull it.

## Troubleshooting Steps

```bash
kubectl describe pod <pod-name>
```

Checked Events for image pull failures.

Verified:

* Image repository
* Image tag
* IAM permissions
* Repository existence

## Resolution

Pushed the correct image to Amazon ECR and confirmed the deployment referenced the correct image tag.

## Learning

Always verify both the image repository and the exact tag before investigating Kubernetes.

---

# 4. Amazon EKS Cluster Created Without Worker Nodes

## Problem

```text
kubectl get nodes

No resources found
```

## Root Cause

The control plane was created successfully, but the managed node group failed because the selected EC2 instance type was unavailable.

## Troubleshooting Steps

* Verified cluster status.
* Listed node groups.
* Inspected CloudFormation events.

## Resolution

Changed the worker node instance type to `t3.small` and recreated the node group.

## Learning

A healthy EKS control plane does not guarantee worker nodes are available.

---

# 5. kubectl Could Not Connect to the Cluster

## Problem

```text
The connection to the server 127.0.0.1:6443 was refused
```

## Root Cause

The local kubeconfig was not configured for the newly created cluster.

## Troubleshooting Steps

Verified the current Kubernetes context.

```bash
kubectl config current-context
```

## Resolution

```bash
aws eks update-kubeconfig \
  --region ap-south-2 \
  --name employee-api-cluster
```

## Learning

Cluster creation and kubeconfig configuration are separate tasks.

---

# 6. Argo CD Was Healthy but the Application Was Not Updated

## Problem

Argo CD reported the application as Healthy, but the old application version continued to run.

## Root Cause

The Helm chart referenced a static image tag (`v2`).

Argo CD synchronized successfully because Git had not changed.

## Resolution

Configured the CD pipeline to update the image tag in `helm/values.yaml` with the Git commit SHA after every successful build.

## Learning

GitOps tools react to Git changes, not container registry changes.

---

# 7. GitHub Actions Failed While Pushing Workflow Changes

## Problem

```text
refusing to allow a Personal Access Token to create or update workflow
```

## Root Cause

The Personal Access Token did not have permission to modify workflow files.

## Resolution

Created a Fine-Grained Personal Access Token with GitHub Actions workflow permissions.

Updated Git credentials.

## Learning

Workflow files require additional GitHub permissions compared to regular source code.

---

# 8. Amazon ECR Authentication Failed

## Problem

```text
no basic auth credentials
```

## Root Cause

Docker was not authenticated with Amazon ECR.

## Resolution

Authenticated Docker using AWS CLI.

```bash
aws ecr get-login-password \
| docker login \
--username AWS \
--password-stdin <ACCOUNT_ID>.dkr.ecr.ap-south-2.amazonaws.com
```

## Learning

Authentication tokens for Amazon ECR expire periodically and must be refreshed.

---

# 9. PostgreSQL Service Was Unreachable

## Problem

The application failed to connect to PostgreSQL.

## Troubleshooting Steps

* Verified PostgreSQL pod status.
* Checked deployment logs.
* Confirmed Kubernetes Service.
* Validated ConfigMap values.
* Tested DNS resolution inside the cluster.

## Resolution

Restarted the PostgreSQL Deployment and confirmed the Service correctly exposed the database.

## Learning

Always verify Pod → Service → DNS → Application connectivity rather than assuming the database is unavailable.

---

# Key DevOps Lessons

* Read Kubernetes Events before modifying manifests.
* Verify CloudFormation when EKS resources fail to provision.
* Use immutable Docker image tags instead of reusing tags.
* Build multi-platform Docker images for heterogeneous environments.
* GitOps deployments depend on Git state, not container registry state.
* Separate CI (build and test) from CD (deployment).
* Use `kubectl describe`, `kubectl logs`, and `kubectl get events` as the first line of troubleshooting.
* Confirm IAM permissions before assuming infrastructure problems.
* Diagnose issues systematically by validating each layer: source code → CI → container image → registry → Kubernetes → application.
