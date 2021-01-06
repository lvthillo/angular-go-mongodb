// This is an environment for testing in my local k8s cluster where the backend is exposed via a NodePort (32001)

export const environment = {
  production: true,
  environmentName: 'k8s',
  gateway: 'http://192.168.50.11:32001',
};