package cmd

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// MARK: Check if the namespace has Pods, Services, and other resources
func inspectNamespace(clientset *kubernetes.Clientset, namespace string) (bool, map[string]int, error) {
	// Initialize a map to store resource counts
	resourceCounts := make(map[string]int)

	// Check for Pods
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false, resourceCounts, err
	}
	resourceCounts["Pods"] = len(pods.Items)

	// Check for Services
	services, err := clientset.CoreV1().Services(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false, resourceCounts, err
	}
	resourceCounts["Services"] = len(services.Items)

	// Check for ConfigMaps
	configMaps, err := clientset.CoreV1().ConfigMaps(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false, resourceCounts, err
	}
	resourceCounts["ConfigMaps"] = len(configMaps.Items)

	// Check for Secrets
	secrets, err := clientset.CoreV1().Secrets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false, resourceCounts, err
	}
	resourceCounts["Secrets"] = len(secrets.Items)

	// Check for Deployments
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false, resourceCounts, err
	}
	resourceCounts["Deployments"] = len(deployments.Items)

	// Check for ReplicaSets
	replicaSets, err := clientset.AppsV1().ReplicaSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false, resourceCounts, err
	}
	resourceCounts["ReplicaSets"] = len(replicaSets.Items)

	// Check for StatefulSets
	statefulSets, err := clientset.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false, resourceCounts, err
	}
	resourceCounts["StatefulSets"] = len(statefulSets.Items)

	// Check for DaemonSets
	daemonSets, err := clientset.AppsV1().DaemonSets(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false, resourceCounts, err
	}
	resourceCounts["DaemonSets"] = len(daemonSets.Items)

	// Check for Jobs
	jobs, err := clientset.BatchV1().Jobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false, resourceCounts, err
	}
	resourceCounts["Jobs"] = len(jobs.Items)

	// Check for CronJobs
	cronJobs, err := clientset.BatchV1().CronJobs(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false, resourceCounts, err
	}
	resourceCounts["CronJobs"] = len(cronJobs.Items)

	// Check for PersistentVolumeClaims
	pvcs, err := clientset.CoreV1().PersistentVolumeClaims(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false, resourceCounts, err
	}
	resourceCounts["PersistentVolumeClaims"] = len(pvcs.Items)

	// Check for Ingresses
	ingresses, err := clientset.NetworkingV1().Ingresses(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return false, resourceCounts, err
	}
	resourceCounts["Ingresses"] = len(ingresses.Items)

    // Check for ServiceAccounts
    serviceAccounts, err := clientset.CoreV1().ServiceAccounts(namespace).List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        return false, resourceCounts, err
    }
    resourceCounts["ServiceAccounts"] = len(serviceAccounts.Items)

    // Check for Roles
    roles, err := clientset.RbacV1().Roles(namespace).List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        return false, resourceCounts, err
    }
    resourceCounts["Roles"] = len(roles.Items)

    // Check for RoleBindings
    roleBindings, err := clientset.RbacV1().RoleBindings(namespace).List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        return false, resourceCounts, err
    }
    resourceCounts["RoleBindings"] = len(roleBindings.Items)

    // Check for NetworkPolicies
    networkPolicies, err := clientset.NetworkingV1().NetworkPolicies(namespace).List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        return false, resourceCounts, err
    }
    resourceCounts["NetworkPolicies"] = len(networkPolicies.Items)

    // Check for ResourceQuotas
    resourceQuotas, err := clientset.CoreV1().ResourceQuotas(namespace).List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        return false, resourceCounts, err
    }
    resourceCounts["ResourceQuotas"] = len(resourceQuotas.Items)

    // Check for LimitRanges
    limitRanges, err := clientset.CoreV1().LimitRanges(namespace).List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        return false, resourceCounts, err
    }
    resourceCounts["LimitRanges"] = len(limitRanges.Items)

    // Check for HorizontalPodAutoscalers
    hpas, err := clientset.AutoscalingV2().HorizontalPodAutoscalers(namespace).List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        return false, resourceCounts, err
    }
    resourceCounts["HorizontalPodAutoscalers"] = len(hpas.Items)

    // Check for PodDisruptionBudgets
    pdbs, err := clientset.PolicyV1().PodDisruptionBudgets(namespace).List(context.TODO(), metav1.ListOptions{})
    if err != nil {
        return false, resourceCounts, err
    }
    resourceCounts["PodDisruptionBudgets"] = len(pdbs.Items)

	// If no Pods and Services are found, the namespace is considered empty of primary resources
	return resourceCounts["Pods"] == 0 && resourceCounts["Services"] == 0 && resourceCounts["ConfigMap"] == 0, resourceCounts, nil
}
