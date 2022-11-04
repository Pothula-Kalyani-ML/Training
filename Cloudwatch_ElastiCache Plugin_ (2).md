## Cloudwatch Elastic Cache Plugin:
This plugin captures metrices for a specified Redis or Memcached cluster or Aggregated cluster metrics.
#### ElastiCache:
ElastiCache is the distributed in memory cache environments in the AWS cloud. 
Amazon ElastiCache supports the Memcached and Redis cache engines.
#### ElastiCache metrics monitoring using Cloudwatch:
Amazon ElastiCache is integrated with Amazon CloudWatch, and automatically publishes 3 modes of Monitoring.
1.	Cluster Level Monitoring
2.	Node Level Monitoring
3.	Aggregated Cluster Monitoring (ElastiCache Monitoring)

#### Cluster level Monitoring:
As an arguement we need to provide following parameters:
**Namespace**: AWS/ElastiCache
**Dimensions**: Each cluster level metric has the following dimensions:
**CacheClusterId**: The ID of the Redis or Memcached for which you want to get metrics.

#### Node level Monitoring:
As an arguement we need to provide following parameters:
**Namespace**: AWS/ElastiCache
**Dimensions**: Each cluster level metric has the following dimensions:
**CacheClusterId**: The ID of the Redis or Memcached for which you want to get metrics
**CacheNodeId**: The Node ID of the Cluster we for which you want to get metrics.

#### Aggregated ElastiCache Metrics:
As an arguement we need to provide following parameters:
**Namespace**: AWS/ElastiCache
**Dimensions**: Empty Dimension

Aggregated ElastiCache Metrics give the Aggregated results of all the cluster metrics of ElastiCache.

### The following are the metrics that we get from cloudwatch

##### Aggregated ElastiCache Metrics:

The AWS/ElastiCache namespace includes the following host-level metrics for individual cache nodes.

| Metric | Unit |
| ------ | ------ |
| CPUCreditBalance | 	Credits (vCPU-minutes) | 
| CPUCreditUsage | 	Credits (vCPU-minutes) | 
| FreeableMemory | 	Bytes | 
| NetworkBytesIn | 	Bytes | 
| NetworkBytesOut | 	Bytes | 
| NetworkPacketsIn | 	Count | 
| NetworkPacketsOut | 	Count | 
| NetworkBandwidthInAllowanceExceeded | 	Count | 
| NetworkConntrackAllowanceExceeded | 	Count | 
| NetworkLinkLocalAllowanceExceeded | 	Count | 
| NetworkBandwidthOutAllowanceExceeded | 	Count | 
| NetworkPacketsPerSecondAllowanceExceeded | 	Count | 
| SwapUsage | 	Bytes | 


