# Exam Topics Questions

@thatonecodes

## Exam KCNA topic 1 question 1 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 1
Topic #: 1

[All KCNA Questions]

What native runtime is Open Container Initiative (OCI) compliant? 
Suggested Answer: A 🗳️ 

A. runC

B. runV

C. kata-containers

D. gvisor

**Answer: A**

**Timestamp: Oct. 3, 2023, 2:29 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/122253-exam-kcna-topic-1-question-1-discussion/)

Comments: zauroc 3 months, 2 weeks ago Selected Answer: A Correct upvoted 1 times ... Suya 1 year, 4 months ago Selected Answer: A runC is the native runtime that is Open Container Initiative (OCI) compliant. upvoted 2 times ... JBangura 1 year, 9 months ago Selected Answer: A A is correct upvoted 2 times ... vishau 2 years, 2 months ago Selected Answer: A this is correct upvoted 2 times ... nvtienanh 2 years, 3 months ago Selected Answer: A Answer is A upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 2 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 2
Topic #: 1

[All KCNA Questions]

Which API object is the recommended way to run a scalable, stateless application on your cluster? 
Suggested Answer: B 🗳️ 

A. ReplicaSet

B. Deployment

C. DaemonSet

D. Pod

**Answer: B**

**Timestamp: Oct. 3, 2023, 2:30 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/122254-exam-kcna-topic-1-question-2-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B Deployment is for stateless applications upvoted 1 times ... Saireddybonthu 1 year, 4 months ago ReplicaSet for statefull applications Deployment is for stateless applications DaemonSet is for application that has to be on every node. Pod is created as part of replicaset, deployment, daemonset upvoted 4 times ... JBangura 1 year, 9 months ago Selected Answer: B B is correct upvoted 2 times ... Anouar_Naggaz 2 years, 2 months ago Deployment upvoted 1 times ... nvtienanh 2 years, 3 months ago Selected Answer: B B is correct upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 3 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 3
Topic #: 1

[All KCNA Questions]

A CronJob is scheduled to run by a user every one hour. What happens in the cluster when it's time for this CronJob to run? 
Suggested Answer: D 🗳️ 

A. Kubelet watches API Server for CronJob objects. When it's time for a Job to run, it runs the Pod directly.

B. Kube-scheduler watches API Server for CronJob objects, and this is why it's called kube-scheduler.

C. CronJob controller component creates a Pod and waits until it finishes to run.

D. CronJob controller component creates a Job. Then the Job controller creates a Pod and waits until it finishes to run.

**Answer: D**

**Timestamp: March 22, 2024, 3:33 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/136898-exam-kcna-topic-1-question-3-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D The CronJob controller is responsible for managing CronJobs and schedules Jobs based on the CronJob's schedule upvoted 1 times ... Saireddybonthu 1 year, 4 months ago CronJob Controller: The CronJob controller is responsible for managing CronJobs. It schedules Jobs based on the CronJob's schedule (e.g., every hour). Job Creation: When it's time for a CronJob to run, the CronJob controller creates a Job resource. This Job represents the workload that needs to be executed. Job Controller: Once the Job is created, the Job controller is responsible for managing this Job. It ensures that the Pods specified by the Job are created and completed successfully. Pod Execution: The Job controller creates one or more Pods based on the Job specification. These Pods execute the tasks defined by the Job. The Job controller waits until the Pod(s) complete their execution. upvoted 4 times ... JBangura 1 year, 9 months ago Selected Answer: D Correct Answer: D upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 4 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 4
Topic #: 1

[All KCNA Questions]

What is the purpose of the kubelet component within a Kubernetes cluster? 
Suggested Answer: D 🗳️ 

A. A dashboard for Kubernetes Clusters that allows management and troubleshooting of applications.

B. A network proxy that runs on each node in your cluster, implementing part of the Kubernetes Service concept.

C. A component that watches for newly created Pods with no assigned node, and selects a node for them to run on.

D. An agent that runs on each node in the cluster. It makes sure that containers are running in a Pod.

**Answer: D**

**Timestamp: Oct. 3, 2023, 2:36 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/122256-exam-kcna-topic-1-question-4-discussion/)

Comments: JBangura 1 year, 3 months ago Selected Answer: D Correct Answer: D upvoted 2 times ... nvtienanh 1 year, 9 months ago Selected Answer: D The given answer is correct upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 5 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 5
Topic #: 1

[All KCNA Questions]

What is the default value for authorization-mode in Kubernetes API server? 
Suggested Answer: B 🗳️ 

A. --authorization-mode=RBAC

B. --authorization-mode=AlwaysAllow

C. --authorization-mode=AlwaysDeny

D. --authorization-mode=ABAC

**Answer: B**

**Timestamp: Aug. 28, 2023, 2:52 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/119231-exam-kcna-topic-1-question-5-discussion/)

Comments: jocasmen94 10 months, 1 week ago Selected Answer: B The answer is B, is always allow by default upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: B This means that by default, the API server allows all requests. However, for production environments, it is recommended to use more restrictive authorization modes such as RBAC (Role-Based Access Control) or Node. upvoted 3 times ... nvtienanh 1 year, 4 months ago Selected Answer: B B upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 6 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 6
Topic #: 1

[All KCNA Questions]

Let's assume that an organization needs to process large amounts of data in bursts, on a cloud-based Kubernetes cluster. For instance: each Monday morning, they need to run a batch of 1000 compute jobs of 1 hour each, and these jobs must be completed by Monday night. What's going to be the most cost-effective method? 
Suggested Answer: B 🗳️ 

A. Run a group of nodes with the exact required size to complete the batch on time, and use a combination of taints, tolerations, and nodeSelectors to reserve these nodes to the batch jobs.

B. Leverage the Kubernetes Cluster Autoscaler to automatically start and stop nodes as they're needed.

C. Commit to a specific level of spending to get discounted prices (with e.g. “reserved instances” or similar mechanisms).

D. Use PriorityСlasses so that the weekly batch job gets priority over other workloads running on the cluster, and can be completed on time.

**Answer: B**

**Timestamp: Aug. 28, 2023, 2:56 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/119233-exam-kcna-topic-1-question-6-discussion/)

Comments: pablokoba Highly Voted 1 year, 2 months ago The most cost-effective method would likely be option B, leveraging the Kubernetes Cluster Autoscaler to automatically start and stop nodes as they're needed. Here's why: Burst processing workloads, like the one described, benefit from the elasticity provided by cloud-based Kubernetes clusters. With Kubernetes Cluster Autoscaler, you can scale your cluster up when there's a demand for more resources (e.g., Monday mornings when the batch jobs need to run) and scale it down during periods of low demand (e.g., after the batch jobs are completed). This ensures that you're only paying for the resources you actually need, avoiding over-provisioning and reducing costs. upvoted 9 times ... randomWaok2 Most Recent 5 months ago Selected Answer: B B is right, Scale when you need: on demand upvoted 1 times ... Mavenh 7 months, 3 weeks ago Selected Answer: B B is right not C cause, C is for continuous running workloads. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: B Cluster Autoscaling: Utilize Kubernetes Cluster Autoscaler to automatically adjust the size of the cluster based on the current workload. This ensures that you only pay for the resources you need when you need them. upvoted 2 times ... yoyo2424 11 months, 2 weeks ago Selected Answer: B B ...Leverage the Kubernetes Cluster Autoscaler to automatically start and stop nodes as they're needed. Explanation: When an organization needs to process large amounts of data in bursts, as described in this scenario, the most cost-effective method is to scale the infrastructure dynamically. The Kubernetes Cluster Autoscaler is specifically designed for this purpose. upvoted 2 times ... SahandFN 1 year ago Selected Answer: C Reserved should be the answer upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 7 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 7
Topic #: 1

[All KCNA Questions]

What is a Kubernetes service with no cluster IP address called? 
Suggested Answer: A 🗳️ 

A. Headless Service

B. Nodeless Service

C. IPLess Service

D. Specless Service

**Answer: A**

**Timestamp: Aug. 28, 2023, 3:02 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/119235-exam-kcna-topic-1-question-7-discussion/)

Comments: jocasmen94 10 months, 1 week ago Selected Answer: A Answer A upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: A A Kubernetes service with no cluster IP address is called a headless service. Headless Service A headless service is created by setting the clusterIP field to None. upvoted 1 times ... EzBL 1 year, 3 months ago A. Headless Service Here's why: Headless Service: This type of service doesn't have a cluster IP address assigned. Instead, it maps directly to the pods it manages, allowing communication with individual pods using their unique IP addresses. Nodeless Service: This term isn't commonly used in Kubernetes. Services can exist and function even without a dedicated node running them. IPLess Service: While this might seem like a logical term, "Headless Service" is the official designation in Kubernetes. Specless Service: A service definition requires a specification file (usually YAML). A service without a spec wouldn't be a valid service at all upvoted 4 times ...
----------------------------------------

## Exam KCNA topic 1 question 8 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 8
Topic #: 1

[All KCNA Questions]

CI/CD stands for: 
Suggested Answer: D 🗳️ 

A. Continuous Information / Continuous Development

B. Continuous Integration / Continuous Development

C. Cloud Integration / Cloud Development

D. Continuous Integration / Continuous Deployment

**Answer: D**

**Timestamp: Oct. 3, 2023, 2:41 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/122258-exam-kcna-topic-1-question-8-discussion/)

Comments: jocasmen94 10 months, 1 week ago Selected Answer: D Answer is D. upvoted 1 times ... dadaarce 1 year, 3 months ago Selected Answer: D I thought B & D are the same. Need to take note to answers like this. Answer is D. upvoted 2 times ... r0xer 1 year, 4 months ago Selected Answer: D The correct answer is D upvoted 2 times ... nvtienanh 1 year, 9 months ago Selected Answer: D D is correct upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 9 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 9
Topic #: 1

[All KCNA Questions]

What default level of protection is applied to the data in Secrets in the Kubernetes API? 
Suggested Answer: D 🗳️ 

A. The values use AES Symmetric Encryption

B. The values are stored in plain text

C. The values are encoded with SHA256 hashes

D. The values are base64 encoded

**Answer: D**

**Timestamp: Aug. 10, 2023, 1:34 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/117703-exam-kcna-topic-1-question-9-discussion/)

Comments: contributor 1 month, 1 week ago Selected Answer: D D is absolutely correct upvoted 1 times ... Diegoflop 7 months, 3 weeks ago Selected Answer: D Answer is D. https://kubernetes.io/docs/concepts/configuration/secret/#working-with-secrets upvoted 1 times ... Diegoflop 7 months, 3 weeks ago Selected Answer: D Answer is D. upvoted 1 times ... jocasmen94 10 months, 1 week ago Selected Answer: D Answer is D. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: D By default, Kubernetes Secrets are base64 encoded. upvoted 1 times ... dadaarce 1 year, 3 months ago Selected Answer: D Kubernetes Secrets store sensitive information such as passwords, OAuth tokens, and SSH keys. The data stored in Secrets is base64 encoded by default. However, it's important to note that base64 encoding is not encryption—it's a reversible encoding scheme. While base64 encoding provides a basic level of obfuscation, it does not provide strong security against unauthorized access. upvoted 4 times ... EzBL 1 year, 3 months ago In Kubernetes, Secrets are stored as base64-encoded strings within etcd, the key-value store used by Kubernetes. Base64 encoding is a method of encoding binary data into ASCII characters, but it is not a form of encryption. Therefore, while base64 encoding obfuscates the data, it does not provide encryption or protection against unauthorized access. It's essential to use additional measures like RBAC (Role-Based Access Control) or encryption mechanisms like encryption at rest to enhance the security of Secrets in Kubernetes. upvoted 2 times ... supersquax 1 year, 4 months ago Selected Answer: D D is correct upvoted 2 times ... nvtienanh 1 year, 10 months ago Selected Answer: D D is correct upvoted 2 times ... eduarte 1 year, 11 months ago Selected Answer: D D is correct upvoted 2 times ... mazabel 1 year, 11 months ago D is correct The secret data is represented as based64-encoded information upvoted 4 times ...
----------------------------------------

## Exam KCNA topic 1 question 10 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 10
Topic #: 1

[All KCNA Questions]

What function does kube-proxy provide to a cluster? 
Suggested Answer: B 🗳️ 

A. Implementing the Ingress resource type for application traffic.

B. Forwarding data to the correct endpoints for Services.

C. Managing data egress from the cluster nodes to the network.

D. Managing access to the Kubernetes API.

**Answer: B**

**Timestamp: Aug. 15, 2023, 11:19 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/118203-exam-kcna-topic-1-question-10-discussion/)

Comments: contributor 1 month, 1 week ago Selected Answer: B B is the answer kube-proxy runs on each node and is responsible for implementing Kubernetes Services. It maintains the network rules that allow Pods to communicate with each other and ensures that traffic sent to a Service’s ClusterIP, NodePort, or LoadBalancer is forwarded to the appropriate Pod endpoints. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: B When a Service is created, kube-proxy sets up the necessary network rules to route traffic to the Pods that are part of that Service. For example, if you have a Service that targets Pods with a specific label, kube-proxy ensures that traffic to the Service's IP address is forwarded to one of the matching Pods upvoted 2 times ... dadaarce 1 year, 3 months ago Selected Answer: B kube-proxy is a network proxy that runs on each node in the Kubernetes cluster. Its primary function is to maintain network rules on each node to forward traffic to the appropriate pods or services based on IP address and port number. upvoted 3 times ... eduarte 1 year, 11 months ago Selected Answer: B B is correct upvoted 4 times ...
----------------------------------------

## Exam KCNA topic 1 question 11 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 11
Topic #: 1

[All KCNA Questions]

How long should a stable API element in Kubernetes be supported (at minimum) after deprecation? 
Suggested Answer: C 🗳️ 

A. 9 months

B. 24 months

C. 12 months

D. 6 months

**Answer: C**

**Timestamp: Aug. 15, 2023, 11:35 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/118205-exam-kcna-topic-1-question-11-discussion/)

Comments: contributor 1 month, 1 week ago Selected Answer: C https://kubernetes.io/docs/reference/using-api/deprecation-policy/?utm_source upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: C A stable API element in Kubernetes should be supported for at least one year after deprecation. This policy ensures that users have sufficient time to migrate to newer versions of the API before the deprecated element is removed in a future release. upvoted 1 times ... pulsefire 1 year, 4 months ago this old question no longer valid: https://kubernetes.io/docs/reference/using-api/deprecation-policy/#:~:text=Rule%20%231%3A%20API,regardless%20of%20track. upvoted 1 times ... omerco61 1 year, 5 months ago Selected Answer: C Rule #7: Deprecated behaviors must function for no less than 1 year after their announced deprecation. https://kubernetes.io/docs/reference/using-api/deprecation-policy/#:~:text=Rule%20%237%3A%20Deprecated%20behaviors%20must%20function%20for%20no%20less%20than%201%20year%20after%20their%20announced%20deprecation. upvoted 2 times ... Charliesco 1 year, 5 months ago Rule #11b: Metrics, after their announced deprecation, must function for no less than: STABLE: 3 releases or 9 months (whichever is longer) BETA: 1 releases or 4 months (whichever is longer) ALPHA: 0 releases Refer to : https://kubernetes.io/docs/reference/using-api/deprecation-policy/#:~:text=STABLE%3A%203%20releases%20or%209%20months%20(whichever%20is%20longer) upvoted 1 times ... eduarte 1 year, 11 months ago Selected Answer: C C is correct upvoted 4 times ...
----------------------------------------

## Exam KCNA topic 1 question 12 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 12
Topic #: 1

[All KCNA Questions]

What is the name of the lightweight Kubernetes distribution built for IoT and edge computing? 
Suggested Answer: B 🗳️ 

A. OpenShift

B. k3s

C. RKE

D. k1s

**Answer: B**

**Timestamp: Aug. 15, 2023, 11:35 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/118206-exam-kcna-topic-1-question-12-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B K3s is a fully compliant Kubernetes distribution that is optimized for resource-constrained environments, making it ideal for IoT and edge computing use cases. It is designed to be lightweight, easy to install, and simple to operate. resources upvoted 1 times ... r0xer 1 year, 4 months ago Easy, B is correct upvoted 2 times ... Spam2210561 1 year, 10 months ago Can someone upload all the 60 KCNA Questions in the discussion forum or mail me? upvoted 1 times ... eduarte 1 year, 11 months ago Selected Answer: B B is correct upvoted 3 times Spam2210561 1 year, 10 months ago Hey can u post all the 60 questions in the discussion forum? upvoted 3 times ... ...
----------------------------------------

## Exam KCNA topic 1 question 13 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 13
Topic #: 1

[All KCNA Questions]

Kubernetes ___ allows you to automatically manage the number of nodes in your cluster to meet demand. 
Suggested Answer: B 🗳️ 

A. Node Autoscaler

B. Cluster Autoscaler

C. Horizontal Pod Autoscaler

D. Vertical Pod Autoscaler

**Answer: B**

**Timestamp: Aug. 28, 2023, 3:10 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/119236-exam-kcna-topic-1-question-13-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B Kubernetes Cluster Autoscaler allows you to automatically manage the number of nodes in your cluster to meet demand. upvoted 1 times ... dadaarce 1 year, 3 months ago Selected Answer: B B. Cluster Autoscaler The Cluster Autoscaler automatically manages the number of nodes in your Kubernetes cluster to meet demand. upvoted 2 times ... fercho 1 year, 7 months ago Selected Answer: B https://github.com/kubernetes/autoscaler upvoted 2 times ... nvtienanh 1 year, 10 months ago Selected Answer: B I think B upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 14 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 14
Topic #: 1

[All KCNA Questions]

Which of the following statements is correct concerning Open Policy Agent (OPA)? 
Suggested Answer: B 🗳️ 

A. The policies must be written in Python language.

B. Kubernetes can use it to validate requests and apply policies.

C. Policies can only be tested when published.

D. It cannot be used outside Kubernetes.

**Answer: B**

**Timestamp: Aug. 28, 2023, 3:13 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/119237-exam-kcna-topic-1-question-14-discussion/)

Comments: dadaarce 1 year, 3 months ago Selected Answer: B OPA is a general-purpose policy engine that can be used to enforce policies across various systems, including Kubernetes. In Kubernetes, OPA can be integrated to validate admission requests, apply policies for resource creation, and enforce custom policies for security, compliance, and other requirements. upvoted 3 times ... fercho 1 year, 7 months ago Selected Answer: B https://www.openpolicyagent.org/docs/latest/#:~:text=The%20Open%20Policy%20Agent%20(OPA,policy%20enforcement%20across%20the%20stack. upvoted 2 times ... nvtienanh 1 year, 10 months ago Selected Answer: B B is correct upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 15 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 15
Topic #: 1

[All KCNA Questions]

In a cloud native world, what does the IaC abbreviation stands for? 
Suggested Answer: B 🗳️ 

A. Infrastructure and Code

B. Infrastructure as Code

C. Infrastructure above Code

D. Infrastructure across Code

**Answer: B**

**Timestamp: Aug. 28, 2023, 3:14 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/119238-exam-kcna-topic-1-question-15-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B infrastructure as code upvoted 1 times ... fercho 1 year, 1 month ago Selected Answer: B B is correct upvoted 2 times ... nvtienanh 1 year, 4 months ago Selected Answer: B B is correct upvoted 3 times ...
----------------------------------------

## Exam KCNA topic 1 question 16 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 16
Topic #: 1

[All KCNA Questions]

In which framework do the developers no longer have to deal with capacity, deployments, scaling and fault tolerance, and OS? 
Suggested Answer: D 🗳️ 

A. Docker Swam

B. Kubernetes

C. Mesos

D. Serverless

**Answer: D**

**Timestamp: Oct. 3, 2023, 2:47 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/122259-exam-kcna-topic-1-question-16-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D In the Serverless framework, developers no longer have to deal with capacity, deployments, scaling, fault tolerance, and OS. upvoted 1 times ... fercho 1 year, 1 month ago Selected Answer: D D is correct upvoted 2 times ... nvtienanh 1 year, 3 months ago Selected Answer: D Should be D upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 17 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 17
Topic #: 1

[All KCNA Questions]

Which of the following characteristics is associated with container orchestration? 
Suggested Answer: B 🗳️ 

A. Application message distribution

B. Dynamic scheduling

C. Deploying application JAR files

D. Virtual Machine distribution

**Answer: B**

**Timestamp: Nov. 18, 2023, 2:02 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/126496-exam-kcna-topic-1-question-17-discussion/)

Comments: shahy0 10 months, 1 week ago Selected Answer: B Container orchestration involves the automated management of containerized applications, including their deployment, scaling, networking, and lifecycle management. One of the key characteristics of container orchestration is dynamic scheduling, which refers to the ability to automatically place containers on the most appropriate nodes in a cluster based on resource availability, constraints, and policies. upvoted 2 times ... Train9 1 year, 1 month ago Yes B is correct upvoted 2 times ... fercho 1 year, 1 month ago Selected Answer: B B is correct upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 18 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 18
Topic #: 1

[All KCNA Questions]

Which of the following workload require a headless service while deploying into the namespace? 
Suggested Answer: A 🗳️ 

A. StatefulSet

B. CronJob

C. Deployment

D. DaemonSet

**Answer: A**

**Timestamp: Aug. 28, 2023, 3:17 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/119240-exam-kcna-topic-1-question-18-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A A headless service is typically used in Kubernetes for workloads that require direct access to individual pods without load balancing. upvoted 1 times ... dadaarce 1 year, 3 months ago Selected Answer: A Headless services are not typically required for the other workload types listed: CronJob: Executes Jobs at a scheduled time, but doesn't necessarily require a headless service. Deployment: Manages stateless applications and can use a standard service for load balancing. DaemonSet: Ensures that a copy of a pod runs on each node in the cluster, but doesn't require a headless service for normal operation. upvoted 2 times ... phcunha 1 year, 3 months ago A. StatefulSet Explanation: StatefulSets are used for applications that maintain a persistent state or have a unique identity, such as databases. Each pod in a StatefulSet typically has a unique name, and it's necessary to ensure discovery and communication between these pods consistently, even when they are scaled up or down. An analogy for StatefulSets could be managing a team of employees in an organization. Each employee has a unique name and a specific role. Even as the team grows or shrinks, it's important to maintain consistent communication among them. Therefore, you need a communication system (service) that can reliably locate each employee, regardless of changes in the team. upvoted 3 times ... pulsefire 1 year, 4 months ago A. https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#:~:text=StatefulSets%20currently%20require%20a%20Headless%20Service%20to%20be%20responsible%20for%20the%20network%20identity%20of%20the%20Pods.%20You%20are%20responsible%20for%20creating%20this%20Service. upvoted 3 times ...
----------------------------------------

## Exam KCNA topic 1 question 19 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 19
Topic #: 1

[All KCNA Questions]

What is Helm? 
Suggested Answer: B 🗳️ 

A. An open source dashboard for Kubernetes.

B. A package manager for Kubernetes applications.

C. A custom scheduler for Kubernetes.

D. An end to end testing project for Kubernetes applications.

**Answer: B**

**Timestamp: Aug. 28, 2023, 3:16 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/119239-exam-kcna-topic-1-question-19-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B Helm is a package manager for Kubernetes that helps you define, install, and manage Kubernetes applications. It uses a packaging format called charts, which are collections of files that describe a related set of Kubernetes resources. upvoted 1 times ... phcunha 1 year, 3 months ago B. A package manager for Kubernetes applications. Explanation: Helm is a package manager for Kubernetes applications. It simplifies the process of deploying, managing, and upgrading complex Kubernetes applications through the use of charts, which are packages of pre-configured Kubernetes resources. Helm allows users to define, install, and upgrade Kubernetes applications with ease, managing dependencies and configurations in a standardized manner. An analogy for Helm could be a delivery service for building projects. Instead of manually gathering all the necessary materials and tools for each project and setting them up individually, Helm acts as a package manager that organizes everything into convenient packages (charts) that can be easily delivered and installed wherever needed. upvoted 1 times ... nvtienanh 1 year, 10 months ago Selected Answer: B B is correct upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 20 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 20
Topic #: 1

[All KCNA Questions]

Which is the correct kubectl command to display logs in real time? 
Suggested Answer: D 🗳️ 

A. kubectl logs -p test-container-1

B. kubectl logs -c test-container-1

C. kubectl logs -l test-container-1

D. kubectl logs -f test-container-1

**Answer: D**

**Timestamp: Aug. 28, 2023, 3:18 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/119241-exam-kcna-topic-1-question-20-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: D The correct kubectl command to display logs in real time from a Kubernetes pod is by using the -f or --follow flag with kubectl logs. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: D The -f or --follow flag is used to stream the logs in real time. If you want to specify a particular container within a pod, you can use the -c flag: upvoted 1 times ... dadaarce 1 year, 3 months ago Selected Answer: D The -f or --follow flag is used to stream logs in real time. upvoted 2 times ... phcunha 1 year, 3 months ago D. kubectl logs -f test-container-1 Explanation: The -f flag in the kubectl logs command stands for "follow," which allows you to continuously stream the logs in real-time as new log entries are added. This is useful for monitoring applications or troubleshooting issues as they occur. An analogy for this command could be tuning in to a live radio broadcast. When you listen to a live radio show, you want to hear the content as it's being broadcasted, rather than just a recording of past broadcasts. Similarly, using kubectl logs -f lets you "tune in" to the ongoing activity of a container, receiving log updates in real-time. upvoted 1 times ... nvtienanh 1 year, 9 months ago Selected Answer: D D is correct upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 21 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 21
Topic #: 1

[All KCNA Questions]

How to load and generate data required before the Pod startup? 
Suggested Answer: A 🗳️ 

A. Use an init container with shared file storage.

B. Use a PVC volume.

C. Use a sidecar container with shared volume.

D. Use another pod with a PVC.

**Answer: A**

**Timestamp: Aug. 28, 2023, 3:19 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/119242-exam-kcna-topic-1-question-21-discussion/)

Comments: phcunha Highly Voted 1 year, 3 months ago A. Use an init container with shared file storage. Explanation: Init containers are executed before the main containers in a Pod are started. They are designed to perform initialization tasks, such as loading and generating data required before the main containers start. Init containers can share file storage with the main containers in the Pod, allowing them to perform tasks like data loading or generation and then making that data available to the main containers. An analogy for using an init container could be preparing ingredients before cooking a meal. Just as you might prepare ingredients like chopping vegetables or marinating meat before starting to cook, init containers prepare data or perform tasks necessary for the main containers to function properly before the main containers start running. upvoted 6 times ... 8378ffe Most Recent 4 months ago Selected Answer: A To load and generate data required before a Kubernetes Pod starts, Init Containers are the primary mechanism. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: A Init Containers are specialized containers that run before the main application containers in a Pod upvoted 1 times ... nvtienanh 1 year, 10 months ago Selected Answer: A A is correct upvoted 3 times ...
----------------------------------------

## Exam KCNA topic 1 question 22 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 22
Topic #: 1

[All KCNA Questions]

What is the core functionality of GitOps tools like Argo CD and Flux? 
Suggested Answer: D 🗳️ 

A. They track production changes made by a human in a Git repository and generate a human-readable audit trail.

B. They replace human operations with an agent that tracks Git commands.

C. They automatically create pull requests when dependencies are outdated.

D. They continuously compare the desired state in Git with the actual production state and notify or act upon differences.

**Answer: D**

**Timestamp: Sept. 18, 2023, 3:53 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/120919-exam-kcna-topic-1-question-22-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D The core functionality of GitOps tools like Argo CD and Flux is to enable continuous delivery and deployment of applications to Kubernetes clusters using Git as the single source of truth. These tools automate the process of synchronizing the desired state of the Kubernetes cluster, as defined in Git repositories, with the actual state of the cluster upvoted 1 times ... phcunha 1 year, 3 months ago An analogy for this functionality could be a thermostat in a room. The desired temperature set on the thermostat represents the configuration defined in Git, while the actual temperature in the room represents the production state. If the actual temperature deviates from the desired temperature, the thermostat (GitOps tool) will automatically adjust the heating or cooling systems to bring the room temperature back to the desired level. Similarly, GitOps tools ensure that the production state of the cluster matches the desired state defined in Git. upvoted 3 times ... phcunha 1 year, 3 months ago D. They continuously compare the desired state in Git with the actual production state and notify or act upon differences. Explanation: GitOps tools like Argo CD and Flux are designed to manage Kubernetes clusters by using Git repositories as the source of truth for cluster configuration. The core functionality of these tools involves continuously comparing the desired state of the cluster, as defined in Git, with the actual state of the cluster running in production. If there are differences between the desired and actual states, these tools will automatically reconcile those differences by either updating the production state to match the desired state or notifying operators of any discrepancies. upvoted 2 times ... nvtienanh 1 year, 9 months ago Selected Answer: D D look right upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 23 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 23
Topic #: 1

[All KCNA Questions]

Which Kubernetes resource workload ensures that all (or some) nodes run a copy of a Pod? 
Suggested Answer: C 🗳️ 

A. ReplicaSet

B. StatefulSet

C. DaemonSet

D. Deployment

**Answer: C**

**Timestamp: Aug. 28, 2023, 3:21 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/119244-exam-kcna-topic-1-question-23-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C A DaemonSet ensures that a copy of a Pod runs on all (or some) nodes in the Kubernetes cluster. upvoted 1 times ... phcunha 1 year, 3 months ago C. DaemonSet Explanation: A DaemonSet ensures that all (or some) nodes run a copy of a Pod. It's typically used for system daemons or cluster services that must run on every node. Each node in the cluster will have exactly one instance of the Pod managed by the DaemonSet. An analogy for a DaemonSet could be streetlights in a city. Just as streetlights are deployed at specific intervals along every street in a city to ensure adequate lighting, DaemonSets ensure that specific Pods are deployed on every node in a Kubernetes cluster to provide essential services or functionalities uniformly across the cluster. upvoted 3 times ... phcunha 1 year, 3 months ago An analogy for this functionality could be a thermostat in a room. The desired temperature set on the thermostat represents the configuration defined in Git, while the actual temperature in the room represents the production state. If the actual temperature deviates from the desired temperature, the thermostat (GitOps tool) will automatically adjust the heating or cooling systems to bring the room temperature back to the desired level. Similarly, GitOps tools ensure that the production state of the cluster matches the desired state defined in Git. upvoted 1 times ... nvtienanh 1 year, 10 months ago Selected Answer: C C is correct upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 24 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 24
Topic #: 1

[All KCNA Questions]

We can extend the Kubernetes API with Kubernetes API Aggregation Layer and CRDs. What is CRD? 
Suggested Answer: A 🗳️ 

A. Custom Resource Definition

B. Custom Restricted Definition

C. Customized RUST Definition

D. Custom RUST Definition

**Answer: A**

**Timestamp: Aug. 28, 2023, 3:31 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/119245-exam-kcna-topic-1-question-24-discussion/)

Comments: plawansai 9 months, 1 week ago Selected Answer: A A is correct upvoted 1 times ... phcunha 1 year, 3 months ago A. Custom Resource Definition Explanation: CRD stands for Custom Resource Definition. It's a Kubernetes extension mechanism that allows users to define their custom resources and their schema, effectively extending the Kubernetes API. Once defined, these custom resources can be managed and interacted with using standard Kubernetes API operations, just like built-in resources like Pods or Deployments. An analogy for CRD could be creating a new type of item in a game. In a video game, developers might introduce custom items that players can collect or use within the game world. These custom items have their unique properties and behaviors, defined by the game developers. Similarly, with CRDs, Kubernetes users can define custom resources with specific properties and behaviors tailored to their application needs. upvoted 4 times ... nvtienanh 1 year, 10 months ago Selected Answer: A A is correct upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 25 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 25
Topic #: 1

[All KCNA Questions]

The Kubernetes project work is carried primarily by SIGs. What does SIG stand for? 
Suggested Answer: A 🗳️ 

A. Special Interest Group

B. Software Installation Guide

C. Support and Information Group

D. Strategy Implementation Group

**Answer: A**

**Timestamp: Aug. 28, 2023, 3:32 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/119246-exam-kcna-topic-1-question-25-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A SIG stands for Special Interest Group. SIGs are responsible for specific areas of the Kubernetes project. They focus on particular aspects of the project, such as networking, storage, or security. upvoted 1 times ... phcunha 1 year, 3 months ago A. Special Interest Group Explanation: SIG stands for Special Interest Group. In the context of Kubernetes, SIGs are groups of contributors focused on specific areas or domains within the project. Each SIG is responsible for overseeing and contributing to the development, maintenance, and improvement of its designated area, such as networking, storage, scalability, or documentation. An analogy for SIGs could be departments within an organization. In a company, different departments (such as marketing, sales, or engineering) focus on specific areas of expertise and collaborate to achieve common goals. Similarly, SIGs in Kubernetes focus on specific aspects of the project and work together to advance the development and functionality of those areas. upvoted 4 times ... EzBL 1 year, 3 months ago Selected Answer: A In the Kubernetes project, SIG stands for Special Interest Group. SIGs are responsible for focused areas of the project and coordinate efforts related to those areas. Each SIG has its own scope, responsibilities, and channels for communication, and they play a crucial role in the development, maintenance, and evolution of Kubernetes. Examples of SIGs include SIG-CLI, SIG-Storage, SIG-Network, etc. upvoted 2 times ... aouihw 1 year, 10 months ago Selected Answer: A Correct https://www.cncf.io/announcements/2019/09/12/cloud-native-computing-foundation-announces-application-delivery-sig/ upvoted 2 times ... nvtienanh 1 year, 10 months ago Selected Answer: A Given answer is correct. upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 26 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 26
Topic #: 1

[All KCNA Questions]

What is the order of 4C’s in Cloud Native Security, starting with the layer that a user has the most control over? 
Suggested Answer: D 🗳️ 

A. Cloud -> Container -> Cluster -> Code

B. Container -> Cluster -> Code -> Cloud

C. Cluster -> Container -> Code -> Cloud

D. Code -> Container -> Cluster -> Cloud

**Answer: D**

**Timestamp: Aug. 28, 2023, 3:33 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/119247-exam-kcna-topic-1-question-26-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: D The 4Cs of Cloud Native Security, ordered from most to least user control, are Code, Container, Cluster, and Cloud. You have the most direct control over your application's Code. You have less direct control over Containers used for deployment, even less over the Cluster (like Kubernetes) that orchestrates them, and the least control over the underlying Cloud provider infrastructure upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: D The 4C’s in Cloud Native Security: Code: Secure coding practices, code reviews, vulnerability scanning. Container: Secure container images, runtime security, least privilege. Cluster: Securing Kubernetes API, RBAC, secure communication. Cloud/Infrastructure: Securing infrastructure, access controls, compliance. upvoted 1 times ... nvtienanh 1 year, 4 months ago Selected Answer: D D make sense upvoted 3 times ...
----------------------------------------

## Exam KCNA topic 1 question 27 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 27
Topic #: 1

[All KCNA Questions]

Which group of container runtimes provides additional sandboxed isolation and elevated security? 
Suggested Answer: C 🗳️ 

A. rune, cgroups

B. docker, containerd

C. runsc, kata

D. crun, cri-o

**Answer: C**

**Timestamp: March 7, 2024, 3:12 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/135364-exam-kcna-topic-1-question-27-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C Kata Containers: Kata Containers combine lightweight virtual machines (VMs) with container technology to provide stronger isolation. Each container runs inside its own VM, offering enhanced security through hardware virtualization upvoted 1 times ... phcunha 1 year, 3 months ago An analogy for this could be security checkpoints in a high-security facility. Just as additional security measures like checkpoints and barriers enhance security within a facility, runsc and Kata Containers add extra layers of isolation and security to containerized environments, ensuring that even if one container is compromised, it doesn't affect the security of other containers or the host system. upvoted 1 times ... phcunha 1 year, 3 months ago runsc (gVisor): It's a lightweight container runtime that runs containers inside a sandboxed environment, providing an additional layer of isolation using user-space kernel emulation. This allows containers to have their own isolated kernel without the overhead of full virtualization. Kata Containers: It's an open-source project that combines the security of virtual machines with the speed and manageability of containers. It uses lightweight VMs to run each container, providing strong isolation between containers without the performance overhead of traditional virtual machines. upvoted 1 times ... phcunha 1 year, 3 months ago C. runsc, kata Explanation: Runsc (gVisor) and Kata Containers are container runtimes that provide additional sandboxed isolation and elevated security compared to traditional container runtimes like Docker or containerd. upvoted 2 times ... pulsefire 1 year, 4 months ago Selected Answer: C C. https://docs.openshift.com/container-platform/4.8/sandboxed_containers/understanding-sandboxed-containers.html#:~:text=OpenShift%20sandboxed%20containers%20support,containment%20through%20VM%20boundaries. upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 28 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 28
Topic #: 1

[All KCNA Questions]

What is the common standard for Service Meshes? 
Suggested Answer: C 🗳️ 

A. Service Mesh Specification (SMS)

B. Service Mesh Technology (SMT)

C. Service Mesh Interface (SMI)

D. Service Mesh Function (SMF)

**Answer: C**

**Timestamp: March 22, 2024, 7:50 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/136907-exam-kcna-topic-1-question-28-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C The common standard for Service Meshes is the Service Mesh Interface (SMI) upvoted 1 times ... phcunha 1 year, 3 months ago C. Service Mesh Interface (SMI) Explanation: The Service Mesh Interface (SMI) is a specification for implementing service mesh functionality across different service mesh implementations. It provides a standard set of APIs for controlling and observing service mesh behavior, allowing for interoperability between various service mesh solutions. An analogy for SMI could be a universal remote control standard. Just as different electronic devices from different manufacturers can be controlled using a universal remote that adheres to a common standard, service mesh implementations can communicate and interoperate effectively using the Service Mesh Interface (SMI) specification. upvoted 4 times ...
----------------------------------------

## Exam KCNA topic 1 question 29 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 29
Topic #: 1

[All KCNA Questions]

Which statement about Ingress is correct? 
Suggested Answer: D 🗳️ 

A. Ingress provides a simple way to track network endpoints within a cluster.

B. Ingress is a Service type like NodePort and ClusterIP.

C. Ingress is a construct that allows you to specify how a Pod is allowed to communicate.

D. Ingress exposes routes from outside the cluster to services in the cluster.

**Answer: D**

**Timestamp: Dec. 17, 2023, 8:03 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/128780-exam-kcna-topic-1-question-29-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D In Kubernetes, an Ingress is an API object that manages external access to services within a cluster, typically HTTP and HTTPS. It provides a way to define rules for routing traffic from outside the cluster to the appropriate services inside the cluster. upvoted 1 times ... phcunha 1 year, 3 months ago D. Ingress exposes routes from outside the cluster to services in the cluster. Explanation: Ingress in Kubernetes is an API object that manages external access to services within a cluster. It provides HTTP and HTTPS routing to services based on incoming requests' hostnames, paths, or other criteria. Ingress exposes routes from outside the cluster to services inside the cluster, acting as an entry point for external traffic to reach services. An analogy for Ingress could be a building's main entrance gate. Just as a main entrance gate controls access to different sections or floors within a building, Ingress controls access to different services within a Kubernetes cluster based on defined routing rules. upvoted 2 times ... AbhishekJoshi 1 year, 6 months ago Selected Answer: D https://kubernetes.io/docs/concepts/services-networking/ingress/#:~:text=Ingress%20exposes%20HTTP%20and%20HTTPS,defined%20on%20the%20Ingress%20resource.&text=An%20Ingress%20may%20be%20configured,offer%20name%2Dbased%20virtual%20hosting. upvoted 3 times ...
----------------------------------------

## Exam KCNA topic 1 question 30 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 30
Topic #: 1

[All KCNA Questions]

What best describes cloud native service discovery? 
Suggested Answer: A 🗳️ 

A. It's a mechanism for applications and microservices to locate each other on a network.

B. It's a procedure for discovering a MAC address, associated with a given IP address.

C. It's used for automatically assigning IP addresses to devices connected to the network.

D. It's a protocol that turns human-readable domain names into IP addresses on the Internet.

**Answer: A**

**Timestamp: March 22, 2024, 7:53 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/136908-exam-kcna-topic-1-question-30-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A Cloud native service discovery is best described as the process by which services in a cloud-native environment automatically discover and communicate with each other without requiring manual configuration. These features enable services to discover and communicate with each other efficiently in dynamic and scalable cloud-native environments. upvoted 1 times ... phcunha 1 year, 3 months ago A. It's a mechanism for applications and microservices to locate each other on a network. Explanation: Cloud native service discovery is a crucial aspect of distributed systems architecture, particularly in microservices environments. It enables applications and microservices to dynamically discover and communicate with each other without hardcoding IP addresses or relying on static configurations. This mechanism allows services to be scalable, resilient, and loosely coupled, facilitating efficient communication and interaction within the cloud environment. An analogy for service discovery could be a directory in a large office building. Just as a directory helps individuals find the locations of various offices or departments within a building, service discovery enables components within a distributed system to locate and communicate with each other dynamically, regardless of their specific network locations or configurations. upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 31 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 31
Topic #: 1

[All KCNA Questions]

What components are common in a service mesh? 
Suggested Answer: D 🗳️ 

A. tracing and log storage

B. circuit breaking and Pod scheduling

C. data plane and runtime plane

D. service proxy and control plane

**Answer: D**

**Timestamp: Dec. 15, 2023, 9:36 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/128626-exam-kcna-topic-1-question-31-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D A service mesh typically includes the following common components: Data Plane: Sidecar proxies for traffic management. Control Plane: Configuration management, service discovery, policy management, and telemetry. Security: mTLS, authentication, and authorization. Traffic Management: Routing, load balancing, retries, and timeouts. upvoted 1 times ... EzBL 1 year, 3 months ago Selected Answer: D In a service mesh architecture, service proxies (such as Envoy, Linkerd proxy, or Istio sidecar) are deployed alongside each service instance to handle communication between services. The control plane consists of various components responsible for configuring, managing, and monitoring the behavior of the service proxies, including features like traffic routing, load balancing, encryption, authentication, and observability. Together, these components enable advanced traffic management, security, and observability capabilities in a distributed microservices environment. upvoted 2 times ... PinkAndBlack 1 year, 7 months ago I can confirn that D is the corrct answer, anyway the question is a bit tricky, because the data plan can be considered as the set of service proxies so even the C answer could be considered upvoted 3 times ...
----------------------------------------

## Exam KCNA topic 1 question 32 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 32
Topic #: 1

[All KCNA Questions]

Which storage operator in Kubernetes can help the system to self-scale, self-heal, etc? 
Suggested Answer: A 🗳️ 

A. Rook

B. Kubernetes

C. Helm

D. Container Storage Interface (CSI)

**Answer: A**

**Timestamp: March 27, 2024, 2:12 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/137317-exam-kcna-topic-1-question-32-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A The storage operator in Kubernetes that can help the system to self-scale, self-heal, and provide other advanced features is Rook. upvoted 1 times ... dadaarce 1 year, 3 months ago Selected Answer: A Rook provides features like self-scaling, self-healing, monitoring, and automation of storage management tasks. It abstracts the complexities of managing distributed storage systems and integrates them seamlessly with Kubernetes, enabling operators to deploy and manage storage clusters as easily as deploying other Kubernetes resources. upvoted 3 times ...
----------------------------------------

## Exam KCNA topic 1 question 33 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 33
Topic #: 1

[All KCNA Questions]

What fields must exist in any Kubernetes object (e.g. YAML) file? 
Suggested Answer: A 🗳️ 

A. apiVersion, kind, metadata

B. kind, namespace, data

C. apiVersion, metadata, namespace

D. kind, metadata, data

**Answer: A**

**Timestamp: March 22, 2024, 7:58 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/136909-exam-kcna-topic-1-question-33-discussion/)

Comments: phcunha Highly Voted 1 year, 3 months ago A. apiVersion, kind, metadata Explanation: In any Kubernetes object definition file, there are three mandatory fields that must exist: apiVersion: Specifies the version of the Kubernetes API that the object uses. This field ensures compatibility and defines the structure of the object. kind: Specifies the type of Kubernetes object being defined, such as Pod, Service, Deployment, etc. This field determines the behavior and functionality of the object. metadata: Contains metadata about the object, such as its name, namespace, labels, and annotations. This metadata is used by Kubernetes to identify and manage the object within the cluster. These fields are essential for Kubernetes to interpret and process the object correctly. The other options do not include all three mandatory fields or contain fields that are not universally required in every Kubernetes object. upvoted 5 times ... 8378ffe Most Recent 4 months ago Selected Answer: A ny Kubernetes object definition, typically expressed in a YAML file, must include the following four top-level fields: apiVersion: Specifies the version of the Kubernetes API that the object belongs to. This determines the schema and available fields for the object. Examples include v1 for core objects like Pods, or apps/v1 for Deployments. kind: Defines the type of Kubernetes object being created. This could be Pod, Deployment, Service, Namespace, etc. metadata: Contains data that uniquely identifies the object and provides administrative information. This typically includes: name: A unique name for the object within its namespace. namespace (optional): The namespace the object belongs to. If not specified, it defaults to the default namespace. labels (optional): Key-value pairs used to organize and select objects. annotations (optional): Unstructured key-value pairs for attaching arbitrary non-identifying metadata. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 34 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 34
Topic #: 1

[All KCNA Questions]

Which of the following would fall under the responsibilities of an SRE? 
Suggested Answer: B 🗳️ 

A. Developing a new application feature.

B. Creating a monitoring baseline for an application.

C. Submitting a budget for running an application in a cloud.

D. Writing policy on how to submit a code change.

**Answer: B**

**Timestamp: March 22, 2024, 8 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/136911-exam-kcna-topic-1-question-34-discussion/)

Comments: phcunha 1 year, 3 months ago Designing strategies for disaster recovery and fault tolerance: Just as the maintenance engineer creates contingency plans for breakdowns or accidents, the SRE devises plans for system failures or outages, ensuring that services can quickly recover and continue functioning. Analyzing system performance metrics: Like reviewing maintenance logs and performance data for each truck, the SRE analyzes metrics from the application's monitoring systems to identify areas for improvement and optimize performance. By considering the SRE's responsibilities in the context of managing a fleet of delivery trucks, we can draw parallels that illustrate the focus on reliability, efficiency, and proactive maintenance inherent in the role. upvoted 1 times ... phcunha 1 year, 3 months ago B. Creating a monitoring baseline for an application. Imagine an SRE as a maintenance engineer for a fleet of delivery trucks. Their primary responsibility is to ensure that the trucks operate reliably and efficiently, delivering goods to customers on time. Here's how the responsibilities of an SRE might align with this analogy: Creating a monitoring baseline for an application: This is akin to installing sensors and monitoring systems on each truck to track parameters like fuel consumption, engine performance, and GPS location. By establishing a baseline for normal truck operation, the SRE can quickly identify deviations or potential issues. Developing and implementing automation for tasks: Similar to designing automated loading and unloading systems for trucks, the SRE automates repetitive tasks in the operation of the fleet, such as scheduling maintenance checks or rerouting trucks in case of road closures. upvoted 3 times ...
----------------------------------------

## Exam KCNA topic 1 question 35 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 35
Topic #: 1

[All KCNA Questions]

What are the initial namespaces that Kubernetes starts with? 
Suggested Answer: A 🗳️ 

A. default, kube-system, kube-public, kube-node-lease

B. default, system, kube-public

C. kube-default, kube-system, kube-main, kube-node-lease

D. kube-default, system, kube-main, kube-primary

**Answer: A**

**Timestamp: Jan. 28, 2024, 5:44 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/132343-exam-kcna-topic-1-question-35-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A B. default, system, kube-public Incorrect. The namespace "system" is not a standard namespace in Kubernetes. C. kube-default, kube-system, kube-main, kube-node-lease Incorrect. The namespaces "kube-default" and "kube-main" are not standard namespaces in Kubernetes. D. kube-default, system, kube-main, kube-primary Incorrect. The namespaces "kube-default", "system", "kube-main", and "kube-primary" are not standard namespaces in Kubernetes. upvoted 2 times ... Error_2k 1 year, 5 months ago correct: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 36 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 36
Topic #: 1

[All KCNA Questions]

What is a probe within Kubernetes? 
Suggested Answer: C 🗳️ 

A. A monitoring mechanism of the Kubernetes API.

B. A pre-operational scope issued by the kubectl agent.

C. A diagnostic performed periodically by the kubelet on a container.

D. A logging mechanism of the Kubernetes API.

**Answer: C**

**Timestamp: March 22, 2024, 8:03 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/136912-exam-kcna-topic-1-question-36-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C Probes are used to determine if a container is running correctly and if it is ready to serve traffic. Kubernetes supports three types of probes: Liveness Probe Readiness Probe Startup Probe upvoted 2 times ... SeaH0rse66 1 year, 1 month ago Selected Answer: C Within Kubernetes, a "probe" refers to a diagnostic mechanism used by the kubelet to check the health of containers running within pods.Probes are configured within the PodSpec of Kubernetes pods and are used to determine if the containers within the pod are healthy and ready to serve traffic. There are three types of probes: Liveness Probe: Determines if the container is still running and healthy. If the liveness probe fails, Kubernetes restarts the container. Readiness Probe: Determines if the container is ready to serve traffic. If the readiness probe fails, the pod is removed from load balancers, and no traffic is routed to it. Startup Probe: Similar to the liveness probe, but only runs during the initial startup of a container. It helps delay the liveness and readiness probes until the application inside the container has started. upvoted 3 times ... phcunha 1 year, 3 months ago Think of a probe in Kubernetes like a heart rate monitor attached to a patient in a hospital. Just as the heart rate monitor continuously checks the patient's heart rate to ensure they are alive and functioning correctly, Kubernetes probes continuously monitor the health of containers to ensure they are running and capable of serving traffic. If the heart rate monitor detects irregularities or a lack of heartbeat, medical staff take action to revive or stabilize the patient. Similarly, if Kubernetes probes detect issues with a container, Kubernetes takes action to restart or replace the container, ensuring the application remains available and responsive. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 37 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 37
Topic #: 1

[All KCNA Questions]

Which Kubernetes feature would you use to guard against split brain scenarios with your distributed application? 
Suggested Answer: D 🗳️ 

A. Replication controllers

B. Consensus protocols

C. Rolling updates

D. StatefulSet

**Answer: D**

**Timestamp: Dec. 22, 2023, 5:24 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/129239-exam-kcna-topic-1-question-37-discussion/)

Comments: Vox 1 month, 1 week ago Selected Answer: D Split brain occurs when multiple replicas of a distributed application (like a database or consensus system) believe they are the leader due to network partitions or miscoordination. • Kubernetes StatefulSets are designed for stateful, distributed applications that require: • Stable network identities (each pod gets a consistent DNS name). • Ordered, graceful deployment and scaling (pods come up one at a time, in sequence). • Persistent storage association (each pod keeps its own volume). • These guarantees help distributed systems maintain quorum and avoid multiple leaders, reducing the risk of split brain upvoted 1 times ... donathon 7 months, 3 weeks ago Selected Answer: B A. Replication controllers: Ensure a specified number of pod replicas are running, but don’t handle split-brain or consensus. C. Rolling updates: Manage application updates with zero downtime, but unrelated to split-brain prevention. D. StatefulSet: Useful for managing stateful applications and preserving identity across restarts, but it doesn’t prevent split-brain conditions. So, B. Consensus protocols is the correct and most relevant answer. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: D A. Replication controllers: Manage the number of replicas of a pod but do not provide the stable network identities or ordered deployment required to prevent split-brain scenarios. B. Consensus protocols: While consensus protocols (like Raft or Paxos) are used within distributed systems to prevent split-brain, they are not a Kubernetes feature. They are implemented within the application itself. C. Rolling updates: Ensure that updates to applications are done gradually to avoid downtime, but they do not address the stable identities or ordered scaling needed to prevent split-brain scenarios. By using StatefulSets, you can effectively manage stateful applications and guard against split-brain scenarios in your distributed applications running on Kubernetes. upvoted 2 times ... 2211094 11 months, 2 weeks ago Selected Answer: B The answer is B upvoted 1 times ... yoyo2424 11 months, 2 weeks ago Selected Answer: B B. Consensus protocols Explanation: Split brain scenarios occur when nodes or components of a distributed application lose communication with each other, leading to conflicting or inconsistent states. To guard against split brain, distributed systems rely on consensus protocols to ensure that a majority of nodes agree on a consistent state, even in the presence of network partitions or failures upvoted 2 times ... Jtrix88772 1 year ago Selected Answer: D Question clearly asks "Which Kubernetes feature". I don't think "Consensus protocols" are part of Kubernetes. upvoted 1 times ... Andrei_Z 1 year, 3 months ago Selected Answer: D The question says: "with your distributed application" so it is D. upvoted 1 times ... EzBL 1 year, 6 months ago Selected Answer: B Consensus protocols, such as those provided by distributed systems like etcd or ZooKeeper, help prevent split-brain scenarios by ensuring that only one leader or primary instance is elected to make decisions or perform critical tasks within the distributed system at any given time. These protocols provide mechanisms for nodes to coordinate and agree on the state of the system, even in the presence of network partitions or failures, thereby mitigating the risk of conflicting or divergent states that could lead to split-brain scenarios. While Kubernetes itself does not provide consensus protocols directly, it often relies on external systems like etcd for managing cluster state and coordination. upvoted 2 times ... SeaH0rse66 1 year, 7 months ago Selected Answer: D D. StatefulSet While consensus protocols are generally used to prevent split brain scenarios in distributed systems, the provided information highlights that StatefulSets are specifically designed to ensure the stability and integrity of distributed and clustered applications. StatefulSets maintain "at most one" semantics, which helps prevent multiple instances of the same identity, reducing the risk of split brain scenarios and data loss in quorum-based systems. StatefulSets are well-suited for applications that require stable network identity and storage, providing mechanisms to manage pod identities, persistent storage, and ordered deployment and scaling. Therefore, StatefulSets are the most appropriate Kubernetes feature to mitigate the risk of split brain scenarios in this context. upvoted 3 times ... hovnival 1 year, 8 months ago Selected Answer: B guys, both copilot and ChatGPT says Consensus protocols. upvoted 2 times SeaH0rse66 1 year, 7 months ago chatGPT is wrong...D. StatefulSet While consensus protocols are generally used to prevent split brain scenarios in distributed systems, the provided information highlights that StatefulSets are specifically designed to ensure the stability and integrity of distributed and clustered applications. StatefulSets maintain "at most one" semantics, which helps prevent multiple instances of the same identity, reducing the risk of split brain scenarios and data loss in quorum-based systems. StatefulSets are well-suited for applications that require stable network identity and storage, providing mechanisms to manage pod identities, persistent storage, and ordered deployment and scaling. Therefore, StatefulSets are the most appropriate Kubernetes feature to mitigate the risk of split brain scenarios in this context. upvoted 1 times ... ... phcunha 1 year, 9 months ago Think of StatefulSet in Kubernetes as a traffic control system at an intersection with multiple lanes. Just as the traffic control system manages the flow of vehicles through the intersection, ensuring that only one lane can proceed at a time to prevent collisions or gridlock, StatefulSet manages the deployment of stateful applications, ensuring that only one instance can be active or "in charge" at any given time to prevent conflicts or inconsistencies in distributed systems. upvoted 1 times ... pulsefire 1 year, 10 months ago Selected Answer: D https://kubernetes.io/docs/tasks/run-application/force-delete-stateful-set-pod/#:~:text=split%20brain%20scenario%20in%20quorum%2Dbased%20systems upvoted 4 times ... sad_schedule 2 years ago https://unofficial-kubernetes.readthedocs.io/en/latest/tasks/manage-stateful-set/upgrade-pet-set-to-stateful-set/#:~:text=StatefulSet%20guards%20against%20split%20brain,brain%20scenarios%20with%20distributed%20applications. upvoted 1 times ... sadsak 2 years ago Shouldn't this be Consensus protocols? upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 38 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 38
Topic #: 1

[All KCNA Questions]

What feature must a CNI support to control specific traffic flows for workloads running in Kubernetes? 
Suggested Answer: D 🗳️ 

A. Border Gateway Protocol

B. IP Address Management

C. Pod Security Policy

D. Network Policies

**Answer: D**

**Timestamp: March 7, 2024, 3:47 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/135366-exam-kcna-topic-1-question-38-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: D Network Policies are a Kubernetes resource that allows you to define rules for how Pods are allowed to communicate with each other and with external network endpoints. For a CNI plugin to enforce these policies, it needs the capability to: upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: D Network Policies are a Kubernetes feature that allows you to control the traffic flow to and from pods. They enable you to define rules for how pods can communicate with each other and with other network endpoints. This is essential for securing workloads and controlling specific traffic flows within a Kubernetes cluster. upvoted 1 times ... phcunha 1 year, 3 months ago Think of Network Policies in Kubernetes as security checkpoints at different entrances to a building. Just as security personnel at each entrance check IDs and verify permissions before allowing individuals to enter specific areas of the building, Network Policies control the flow of network traffic within the Kubernetes cluster, ensuring that only authorized Pods can communicate with each other and that access to sensitive services is restricted according to defined rules. upvoted 2 times ... pulsefire 1 year, 4 months ago Selected Answer: D https://kubernetes.io/docs/concepts/services-networking/network-policies/#:~:text=Network%20Policies-,Network%20Policies,might%20consider%20using%20Kubernetes%20NetworkPolicies%20for%20particular%20applications%20in%20your%20cluster.,-NetworkPolicies%20are%20an upvoted 3 times ...
----------------------------------------

## Exam KCNA topic 1 question 39 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 39
Topic #: 1

[All KCNA Questions]

What is the main role of the Kubernetes DNS within a cluster? 
Suggested Answer: D 🗳️ 

A. Acts as a DNS server for virtual machines that are running outside the cluster.

B. Provides a DNS as a Service, allowing users to create zones and registries for domains that they own.

C. Allows Pods running in dual stack to convert IPv6 calls into IPv4 calls.

D. Provides consistent DNS Names for Pods and Services for workloads that need to communicate with each other.

**Answer: D**

**Timestamp: March 22, 2024, 8:07 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/136913-exam-kcna-topic-1-question-39-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D The main role of Kubernetes DNS within a cluster is to provide consistent DNS names for pods and services, enabling reliable communication between workloads. This service discovery mechanism allows applications and services to locate and communicate with each other using DNS names rather than IP addresses, which can change over time. upvoted 1 times ... SeaH0rse66 1 year, 1 month ago Selected Answer: D Kubernetes DNS provides DNS resolution services within a Kubernetes cluster. It ensures that Pods and Services can be discovered and accessed by other Pods and Services using consistent DNS names. This enables seamless communication between different components within the Kubernetes cluster, regardless of their underlying network configurations. upvoted 2 times ... phcunha 1 year, 3 months ago Think of Kubernetes DNS as the address book in a large office building. Just as the address book provides a consistent reference for finding the location of different departments or individuals within the building, Kubernetes DNS provides consistent DNS names for Pods and Services, allowing workloads within the cluster to discover and communicate with each other efficiently, regardless of their dynamic IP addresses or locations. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 40 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 40
Topic #: 1

[All KCNA Questions]

Scenario: You have a Kubernetes cluster hosted in a public cloud provider. When trying to create a Service of type LoadBalancer, the external-ip is stuck in the "Pending" state. Which Kubernetes component is failing in this scenario? 
Suggested Answer: A 🗳️ 

A. Cloud Controller Manager

B. Load Balancer Manager

C. Cloud Architecture Manager

D. Cloud Load Balancer Manager

**Answer: A**

**Timestamp: Jan. 12, 2024, 3:32 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/130945-exam-kcna-topic-1-question-40-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: A The Cloud Controller Manager is responsible for interacting with the underlying public cloud provider's API to provision and manage cloud resources, including: upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: A The Cloud Controller Manager is a Kubernetes component that integrates with the cloud provider's API to manage resources such as load balancers, nodes, and storage. When you create a Service of type LoadBalancer, the Cloud Controller Manager is responsible for provisioning the external load balancer and assigning an external IP address. upvoted 1 times ... SeaH0rse66 1 year, 1 month ago Selected Answer: A B,C,D are not standards of Kubernetes. The Cloud Controller Manager is a Kubernetes component responsible for managing cloud-specific resources and integrations, including load balancers provided by the cloud provider. It interacts with the cloud provider's API to provision, configure, and manage resources such as load balancers. If the Service's external IP address remains in the "Pending" state, it suggests that there may be an issue with the Cloud Controller Manager's ability to communicate with the cloud provider's API or to provision the necessary resources. upvoted 2 times ... sad_schedule 1 year, 6 months ago https://kubernetes.io/docs/concepts/architecture/cloud-controller/#service-controller upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 41 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 41
Topic #: 1

[All KCNA Questions]

What are the characteristics for building every cloud-native application? 
Suggested Answer: D 🗳️ 

A. Resiliency, Operability, Observability, Availability

B. Resiliency, Containerd, Observability, Agility

C. Kubernetes, Operability, Observability, Availability

D. Resiliency, Agility, Operability, Observability

**Answer: D**

**Timestamp: Jan. 18, 2024, 9:23 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/131468-exam-kcna-topic-1-question-41-discussion/)

Comments: zauroc 3 months, 2 weeks ago Selected Answer: D D is Correct upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: D Resiliency: The ability of an application to recover from failures and continue to function. This includes handling transient faults, ensuring data consistency, and maintaining service availability. Agility: The ability to quickly develop, test, and deploy new features and updates. This often involves practices like continuous integration and continuous deployment (CI/CD), as well as the use of microservices architecture. Operability: The ease with which an application can be managed, monitored, and maintained in production. This includes aspects like automated deployments, scaling, and self-healing capabilities. Observability: The ability to gain insights into the internal state of an application through metrics, logs, and traces. This helps in monitoring the health of the application, diagnosing issues, and improving performance. upvoted 1 times ... 2211094 12 months ago Selected Answer: D The correct answer is D. Resiliency, Agility, Operability, Observability. upvoted 1 times ... z2salman 1 year, 1 month ago Selected Answer: A As stated above, Availability is a key rationality to implement containerization of applications as opposed to agility. upvoted 1 times ... alex78 1 year, 2 months ago Selected Answer: D A cloud native application is engineered to run on a platform and is designed for resiliency, agility, operability, and observability. Resiliency embraces failures instead of trying to prevent them; it takes advantage of the dynamic nature of running on a platform. Agility allows for fast deployments and quick iterations. Operability adds control of application life cycles from inside the application instead of relying on external processes and monitors. Observability provides information to answer questions about application state. https://www.oreilly.com/library/view/cloud-native-infrastructure/9781491984291/ch01.html upvoted 2 times ... alex78 1 year, 2 months ago Selected Answer: D https://www.oreilly.com/library/view/cloud-native-infrastructure/9781491984291/ch01.html upvoted 2 times ... hovnival 1 year, 2 months ago Selected Answer: D Folks I ran the same questions again in Copilot and week after it says correct answer is D. I am very sorry about that Apologies. Confirmed by ChatGPT answer is D. It is a bit frustrating that copilot is not consistent in aswer as I did copy paste like last time. upvoted 2 times ... hovnival 1 year, 2 months ago Selected Answer: A Copilot says A upvoted 1 times ... JBangura 1 year, 3 months ago Selected Answer: D https://lists.cncf.io/g/cncf-toc/message/1636 upvoted 2 times ... EzBL 1 year, 3 months ago Selected Answer: D Availability is highly desirable, but it's a consequence of achieving resiliency through proper design and implementation. upvoted 2 times ... 7a9e5e5 1 year, 3 months ago I am agree with correct response is A. High Availability is so principal to deploy Cloud Native aplications and it is an crucial to keep the aplication without interruptions. Agility is important but not most importan than Availability. upvoted 1 times ... unkun 1 year, 5 months ago It should Be A upvoted 2 times ... majkisermi98 1 year, 5 months ago Selected Answer: A Answer is wrong, it should be A. Availability is a crucial principle, Agility is not upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 42 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 42
Topic #: 1

[All KCNA Questions]

What does CNCF stand for? 
Suggested Answer: B 🗳️ 

A. Cloud Native Community Foundation

B. Cloud Native Computing Foundation

C. Cloud Neutral Computing Foundation

D. Cloud Neutral Community Foundation

**Answer: B**

**Timestamp: Oct. 18, 2024, 7:48 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/149725-exam-kcna-topic-1-question-42-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B Cloud Native Computing Foundation (CNCF) upvoted 1 times ... heeloco 1 year, 2 months ago Selected Answer: B B) upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 43 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 43
Topic #: 1

[All KCNA Questions]

Kubernetes supports multiple virtual clusters backed by the same physical cluster. These virtual clusters are called: 
Suggested Answer: A 🗳️ 

A. namespaces

B. containers

C. hypervisors

D. cgroups

**Answer: A**

**Timestamp: Jan. 31, 2024, 10:19 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/132540-exam-kcna-topic-1-question-43-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A Namespaces in Kubernetes are virtual clusters that provide resource isolation, access control, and organization within a single physical cluster. They enable multiple teams or projects to share the same physical cluster without interfering with each other. upvoted 1 times ... SeaH0rse66 1 year, 1 month ago Selected Answer: A https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/ "Namespaces are intended for use in environments with many users spread across multiple teams, or projects. For clusters with a few to tens of users, you should not need to create or think about namespaces at all. Start using namespaces when you need the features they provide." upvoted 2 times ... phcunha 1 year, 3 months ago Think of namespaces in Kubernetes as different floors in a large office building. Each floor represents a separate workspace for different teams or departments within the organization. Just as each floor has its own set of offices, meeting rooms, and facilities, each namespace in Kubernetes has its own set of resources, configurations, and access controls, providing isolation and organization within the cluster. upvoted 1 times ... unkun 1 year, 5 months ago https://jamesdefabia.github.io/docs/user-guide/namespaces/ Answer is correct "A" upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 44 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 44
Topic #: 1

[All KCNA Questions]

What component enables end users, different parts of the Kubernetes cluster, and external components to communicate with one another? 
Suggested Answer: C 🗳️ 

A. kubectl

B. AWS Management Console

C. Kubernetes API

D. Google Cloud SDK

**Answer: C**

**Timestamp: May 17, 2024, 4:58 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/140799-exam-kcna-topic-1-question-44-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C The Kubernetes API is the central component that enables communication between end users, different parts of the Kubernetes cluster, and external components. It provides a unified interface for interacting with the cluster, allowing users and components to create, update, delete, and retrieve Kubernetes resources. upvoted 1 times ... 2211094 12 months ago Selected Answer: C The correct answer is C. Kubernetes API. The Kubernetes API server is the central component that enables communication between end users, different parts of the Kubernetes cluster, and external components. upvoted 1 times ... SeaH0rse66 1 year, 1 month ago Selected Answer: C The Kubernetes API serves as the primary interface for communication between end users, different parts of the Kubernetes cluster, and external components. upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 45 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 45
Topic #: 1

[All KCNA Questions]

Which command will list the resource types that exist within a cluster? 
Suggested Answer: A 🗳️ 

A. kubectl api-resources

B. kubectl get namespaces

C. kubectl api-versions

D. curl https://kubectrl/namespaces

**Answer: A**

**Timestamp: Sept. 18, 2024, 8:21 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/147797-exam-kcna-topic-1-question-45-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A kubectl api-resources: This command displays a list of all resource types that are available in the Kubernetes API. It includes both built-in resources (like pods, services, and deployments) and custom resources that may be defined by Custom Resource Definitions (CRDs). upvoted 1 times ... vspringe 1 year, 3 months ago Selected Answer: A kubectl api-resources upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 46 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 46
Topic #: 1

[All KCNA Questions]

Which of these components is part of the Kubernetes Control Plane? 
Suggested Answer: B 🗳️ 

A. coredns

B. cloud-controller-manager

C. kube-proxy

D. kubelet

**Answer: B**

**Timestamp: Jan. 31, 2024, 10:25 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/132543-exam-kcna-topic-1-question-46-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B cloud-controller-manager: This component integrates with the cloud provider's API to manage cloud-specific resources like load balancers, storage, and networking. upvoted 1 times ... heeloco 1 year, 2 months ago Selected Answer: B B upvoted 1 times ... unkun 1 year, 11 months ago https://kubernetes.io/docs/concepts/architecture/cloud-controller/ Answer is correct upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 47 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 47
Topic #: 1

[All KCNA Questions]

Which of the following systems is NOT compatible with the CRI runtime interface standard? 
Suggested Answer: C 🗳️ 

A. CRI-0

B. dockershim

C. systemd

D. containerd

**Answer: C**

**Timestamp: Feb. 8, 2024, 9:22 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/133402-exam-kcna-topic-1-question-47-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C systemd is not a container runtime and does not interact with the CRI standard. CRI-0 (CRI-O), dockershim, and containerd are all container runtimes that are compatible with the CRI standard, allowing them to be used with Kubernetes. upvoted 1 times ... hovnival 1 year, 2 months ago Selected Answer: C CRI - Container Runtime Interface. systemd is not container runtime interface. upvoted 2 times ... JBangura 1 year, 3 months ago https://kubernetes.io/docs/tasks/administer-cluster/migrating-from-dockershim/check-if-dockershim-removal-affects-you/ upvoted 1 times ... omerco61 1 year, 5 months ago Selected Answer: C systemd: systemd is not a container runtime. It is an initialization system and service manager for Linux. upvoted 3 times ...
----------------------------------------

## Exam KCNA topic 1 question 48 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 48
Topic #: 1

[All KCNA Questions]

What is a key feature of a container network? 
Suggested Answer: B 🗳️ 

A. Proxying REST requests across a set of containers.

B. Allowing containers running on separate hosts to communicate.

C. Allowing containers on the same host to communicate.

D. Caching remote disk access.

**Answer: B**

**Timestamp: Sept. 18, 2024, 8:28 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/147798-exam-kcna-topic-1-question-48-discussion/)

Comments: zauroc 3 months, 2 weeks ago Selected Answer: B B is correct upvoted 1 times ... donathon 7 months, 3 weeks ago Selected Answer: B A key feature of a container network in Kubernetes (or any container orchestration system) is to enable network communication between containers, even when they are running on different nodes in the cluster. This is essential for distributed applications and microservices to function properly. A. Proxying REST requests: This is typically handled by service meshes or API gateways, not the container network itself. C. Same-host communication: While important, it's a basic feature; cross-host communication is more critical and complex. D. Caching remote disk access: This relates to storage, not networking. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: B A key feature of a container network is to enable communication between containers running on separate hosts. This is essential for distributed applications where different components of the application may be running on different nodes in a cluster. Container networks provide the necessary infrastructure to ensure that containers can discover and communicate with each other across the network, regardless of the physical host they are running on. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: B A key feature of a container network is to enable communication between containers running on separate hosts. This is essential for distributed applications where different components of the application may be running on different nodes in a cluster. Container networks provide the necessary infrastructure to ensure that containers can discover and communicate with each other across the network, regardless of the physical host they are running on. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: C A key feature of a container network is to enable communication between containers running on separate hosts. This is essential for distributed applications where different components of the application may be running on different nodes in a cluster. Container networks provide the necessary infrastructure to ensure that containers can discover and communicate with each other across the network, regardless of the physical host they are running on. upvoted 1 times ... 2211094 12 months ago Selected Answer: C The correct answer is C. systemd. Systemd is not a container runtime and does not implement the Container Runtime Interface (CRI) standard. upvoted 1 times ... vspringe 1 year, 3 months ago Selected Answer: B B - Allowing containers running on separate hosts to communicate upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 49 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 49
Topic #: 1

[All KCNA Questions]

How can you monitor the progress for an updated Deployment/DaemonSets/StatefulSets? 
Suggested Answer: D 🗳️ 

A. kubectl rollout watch

B. kubectl rollout progress

C. kubectl rollout state

D. kubectl rollout status

**Answer: D**

**Timestamp: April 22, 2024, 1:47 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/139360-exam-kcna-topic-1-question-49-discussion/)

Comments: hovnival Highly Voted 1 year, 8 months ago Selected Answer: A In summary, both commands "kubectl rollout status" and "kubectl rollout watch" provide a real-time feedback on rollout progress but "kubectl rollout watch" continuously watches until completion while "kubectl rollout status" provides a snapshot of the current status. upvoted 5 times ... donathon Most Recent 7 months, 3 weeks ago Selected Answer: D kubectl rollout status allows you to monitor the progress of a rollout for: Deployments DaemonSets StatefulSets It shows whether the rollout is complete, still in progress, or if there are any issues. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: D The kubectl rollout status command is used to monitor the progress of an updated Deployment, DaemonSet, or StatefulSet in Kubernetes. It provides real-time updates on the status of the rollout, including whether it is complete or if there are any issues. upvoted 1 times ... Hach233 1 year, 2 months ago Selected Answer: D D upvoted 1 times ... heniabida 1 year, 2 months ago Selected Answer: D ❯ kubectl rollout watch error: unknown command "watch" Available Commands: history View rollout history pause Mark the provided resource as paused restart Restart a resource resume Resume a paused resource status Show the status of the rollout undo Undo a previous rollout --> So answer is D upvoted 2 times ... yoyo2424 1 year, 2 months ago D is correct answer upvoted 1 times ... mc2301 1 year, 3 months ago kubectl rollout watch Purpose: This command is not a standard kubectl command. It seems like a misunderstanding or a mix-up with kubectl get events --watch or similar commands that watch for changes in resources. Behavior: If you meant kubectl get events --watch, it watches for changes in events related to the resources in the cluster, providing real-time updates. Answer: D upvoted 1 times ... miskill 1 year, 3 months ago Selected Answer: D D. is correct "kubectl rollout status" upvoted 1 times ... vspringe 1 year, 3 months ago Selected Answer: D D. kubectl rollout status Correct. kubectl rollout watch is not a valid Kubernetes command. upvoted 1 times ... miskill 1 year, 4 months ago Selected Answer: D The correct answer to the question "How can you monitor the progress for an updated Deployment/DaemonSets/StatefulSets?" is D. kubectl rollout status. Explanation: The kubectl rollout status command is used to monitor the progress of a rollout for Kubernetes resources such as Deployments, DaemonSets, or StatefulSets. This command provides real-time updates on whether the new Pods are being created correctly and whether the old ones are being terminated. It ensures that the rollout is proceeding as expected and alerts if any issues occur during the update. Other options like kubectl rollout watch, kubectl rollout progress, and kubectl rollout state are not valid Kubernetes commands for monitoring rollout progress. The proper command to use for tracking progress is kubectl rollout status upvoted 1 times ... SeaH0rse66 1 year, 7 months ago Selected Answer: D Option A (kubectl rollout watch) is not the correct choice because the kubectl rollout watch command is used to continuously monitor the rollout progress of a deployment, allowing you to watch for changes in the rollout status in real-time. However, it is specifically designed for monitoring deployments and does not directly apply to DaemonSets or StatefulSets. While you could use kubectl rollout watch for deployments, it may not provide accurate or relevant information for DaemonSets or StatefulSets. Therefore, option A is not the most suitable command for monitoring the progress of updates to DaemonSets or StatefulSets. On the other hand, option D (kubectl rollout status) is a more generic command that works for all types of rollouts, including deployments, DaemonSets, and StatefulSets. It provides detailed information about the status of the rollout, making it a better choice for monitoring the progress of updates across different types of workload controllers in Kubernetes. upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 50 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 50
Topic #: 1

[All KCNA Questions]

What is the goal of load balancing? 
Suggested Answer: D 🗳️ 

A. Automatically measure request performance across instances of an application.

B. Automatically distribute requests across different versions of an application.

C. Automatically distribute instances of an application across the cluster.

D. Automatically distribute requests across instances of an application.

**Answer: D**

**Timestamp: Sept. 18, 2024, 8:34 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/147799-exam-kcna-topic-1-question-50-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D The goal of load balancing is to automatically distribute incoming network traffic (requests) across multiple instances of an application. This helps to ensure that no single instance becomes overwhelmed with too many requests, which can improve the overall performance, reliability, and availability of the application. upvoted 1 times ... vspringe 1 year, 3 months ago Selected Answer: D D. Automatically distribute requests across instances of an application. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 51 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 51
Topic #: 1

[All KCNA Questions]

How are ReplicaSets and Deployments related? 
Suggested Answer: A 🗳️ 

A. Deployments manage ReplicaSets and provide declarative updates to Pods.

B. ReplicaSets manage stateful applications, Deployments manage stateless applications.

C. Deployments are runtime instances of ReplicaSets.

D. ReplicaSets are subsets of Jobs and CronJobs which use imperative Deployments.

**Answer: A**

**Timestamp: Aug. 11, 2023, 9:27 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/117929-exam-kcna-topic-1-question-51-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A Deployments and ReplicaSets are closely related in Kubernetes. A Deployment is a higher-level abstraction that manages ReplicaSets and provides declarative updates to Pods. When you create or update a Deployment, Kubernetes automatically creates or updates the underlying ReplicaSet to match the desired state specified in the Deployment. upvoted 1 times ... 2211094 12 months ago Selected Answer: A How are ReplicaSets and Deployments related? A. Deployments manage ReplicaSets and provide declarative updates to Pods. B. ReplicaSets manage stateful applications, Deployments manage stateless applications. C. Deployments are runtime instances of ReplicaSets. D. ReplicaSets are subsets of Jobs and CronJobs which use imperative Deployments. upvoted 1 times ... SeaH0rse66 1 year, 1 month ago Selected Answer: A Deployments are higher-level abstractions in Kubernetes that manage ReplicaSets.ReplicaSets, on the other hand, are lower-level controllers responsible for maintaining a specified number of identical Pods to ensure high availability and fault tolerance. ReplicaSets are used by Deployments to manage the lifecycle of Pods, including scaling, rolling updates, and maintaining a desired number of replicas. upvoted 2 times ... AzureDP900 1 year, 7 months ago A is correct upvoted 2 times ... LoloMaceto 1 year, 7 months ago Selected Answer: A A is correct upvoted 2 times ... nvtienanh 1 year, 9 months ago Selected Answer: A A is correct upvoted 2 times ... mazabel 1 year, 11 months ago The correct is B upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 52 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 52
Topic #: 1

[All KCNA Questions]

What factors influence the Kubernetes scheduler when it places Pods on nodes? 
Suggested Answer: A 🗳️ 

A. Pod memory requests, node taints, and Pod affinity.

B. Pod labels, node labels, and request labels.

C. Node taints, node level, and Pod priority.

D. Pod priority, container command, and node labels.

**Answer: A**

**Timestamp: Jan. 12, 2025, 12:57 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/154426-exam-kcna-topic-1-question-52-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A Pod Memory Requests: The scheduler considers the resource requests (CPU and memory) specified for the Pod to ensure that the node has sufficient resources to accommodate the Pod. Node Taints: Nodes can be tainted to repel certain Pods. The scheduler checks for taints on nodes and tolerations on Pods to determine if a Pod can be scheduled on a particular node. Pod Affinity and Anti-Affinity: Pod affinity rules specify that certain Pods should be placed on the same node or close to each other, while anti-affinity rules specify that certain Pods should not be placed on the same node or close to each other. The scheduler considers these rules when making placement decisions. upvoted 1 times ... 2211094 12 months ago Selected Answer: A A. Pod memory requests, node taints, and Pod affinity. The Kubernetes scheduler takes into account factors like memory and CPU requests/limits, node taints and tolerations, and Pod affinity and anti-affinity rules to decide where to place Pods. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 53 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 53
Topic #: 1

[All KCNA Questions]

What is the core metric type in Prometheus used to represent a single numerical value that can go up and down? 
Suggested Answer: D 🗳️ 

A. Summary

B. Counter

C. Histogram

D. Gauge

**Answer: D**

**Timestamp: April 22, 2024, 1:57 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/139361-exam-kcna-topic-1-question-53-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D The core metric type in Prometheus used to represent a single numerical value that can go up and down is a Gauge. upvoted 1 times ... mc2301 1 year, 3 months ago The Prometheus data model provides four core metrics: Counter: A value that increases, like a request or error count Gauge: Values that increase or decrease, like memory size Histogram: A sample of observations, like request duration or response size Summary: Similar to a histogram, but also provides the total count of observations. Seems like "up and down" is the most exact match for "increase and decrease", hence D upvoted 1 times ... hovnival 1 year, 8 months ago Selected Answer: D Copilot says D. upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 54 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 54
Topic #: 1

[All KCNA Questions]

What is the primary mechanism to identify grouped objects in a Kubernetes cluster? 
Suggested Answer: B 🗳️ 

A. Custom Resources

B. Labels

C. Label Selector

D. Pod

**Answer: C**

**Timestamp: March 7, 2024, 4:57 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/135369-exam-kcna-topic-1-question-54-discussion/)

Comments: zauroc 3 months, 2 weeks ago Selected Answer: C C is correct upvoted 1 times ... 884f7cc 5 months, 4 weeks ago Selected Answer: C 원문 Kubernetes 문서 기준으로 보면: “Via a label selector, the client/user can identify a set of objects. The label selector is the core grouping primitive in Kubernetes.” kubernetes.io +1 reddit.com +1 reddit.com +11 kubernetes.io +11 unofficial-kubernetes.readthedocs.io +11 즉, **라벨 셀렉터(Label Selector)**가 명확히 “그룹화에 가장 핵심적인(주된) 메커니즘”이라고 규정되어 있습니다. upvoted 1 times ... donathon 7 months, 3 weeks ago Selected Answer: B B. Labels Explanation: In Kubernetes, labels are the primary mechanism used to identify and group objects such as Pods, Services, Deployments, etc. 🔖 Labels Key-value pairs attached to objects Used to organize, select, and operate on subsets of objects Breakdown of the other options: A. Custom Resources → ❌ Extend Kubernetes functionality, not used for grouping. C. Label Selector → ❌ Used to query or select objects based on labels, but not the grouping mechanism itself. D. Pod → ❌ A workload object, not a grouping mechanism. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: B Labels are the primary mechanism to identify and group objects in a Kubernetes cluster. Labels are key-value pairs that can be attached to Kubernetes objects such as Pods, Services, and Deployments. They are used to organize and select subsets of objects based on specific criteria. upvoted 1 times ... 2211094 12 months ago Selected Answer: B B. Labels Labels are the primary mechanism in Kubernetes to identify and group objects. They are key/value pairs that can be attached to objects such as Pods, Nodes, and Services to organize and select subsets of objects. upvoted 1 times ... abitwrong 1 year ago Selected Answer: B Labels are key-value pairs that you attach to Kubernetes objects. They define attributes and allow you to "tag" objects to group or identify them. Label Selectors are used to query or filter objects that have certain labels. While selectors allow you to find or match objects based on labels, they don't define the labels themselves. Labels are the mechanism used to mark or identify objects. Without labels, there would be nothing for Label Selectors to query. Label Selectors are only tools to query/filter based on existing labels. upvoted 3 times ... yoyo2424 1 year, 2 months ago correct answer B upvoted 2 times ... miskill 1 year, 3 months ago Selected Answer: B The correct answer is: B. Labels Explanation: Labels are the primary mechanism used to identify and group objects in a Kubernetes cluster. They are key-value pairs attached to objects (such as Pods, Services, etc.) and can be used to organize, select, and filter resources within the cluster. Although C. Label Selector is also used to query and select objects based on labels, the labels themselves (B) are the mechanism that allows objects to be grouped. upvoted 1 times ... EzBL 1 year, 6 months ago Selected Answer: C Labels: Labels are key-value pairs attached to various Kubernetes objects and act as a metadata mechanism. While essential for grouping, they are the attributes used for grouping, not the grouping mechanism itself. Label Selectors, on the other hand, leverage the information provided by labels to create logical groupings: • Filtering based on Labels: They act as queries or expressions that filter objects based on their labels. You can define requirements that specify the label key and value combinations an object must possess to be considered part of the group. • Building Groups: By using label selectors, you can identify sets of objects that share specific characteristics defined by their labels. This essentially forms logical groups of objects within the cluster. upvoted 2 times ... SeaH0rse66 1 year, 7 months ago Selected Answer: B Labels and label selectors work hand in hand to identify and group objects, so it's easy to see why it might seem tricky. Labels provide the metadata for identification and grouping, while label selectors enable users to query and select objects based on those labels. Both are essential components of Kubernetes's flexibility and functionality. upvoted 3 times ... hovnival 1 year, 8 months ago Selected Answer: C Copilot: Label selectors allow you to group and select objects based on labels, making them a fundamental part of Kubernetes resource management. upvoted 2 times ... JBangura 1 year, 9 months ago Selected Answer: C Correct Answer: C upvoted 2 times ... pulsefire 1 year, 10 months ago Selected Answer: C should be C, not B https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#:~:text=Via%20a%20label%20selector%2C%20the%20client/user%20can%20identify%20a%20set%20of%20objects.%20The%20label%20selector%20is%20the%20core%20grouping%20primitive%20in%20Kubernetes. upvoted 2 times ... pulsefire 1 year, 10 months ago Selected Answer: B should be C, not B https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#:~:text=Via%20a%20label%20selector%2C%20the%20client/user%20can%20identify%20a%20set%20of%20objects.%20The%20label%20selector%20is%20the%20core%20grouping%20primitive%20in%20Kubernetes. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 55 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 55
Topic #: 1

[All KCNA Questions]

What is the name of the Kubernetes resource used to expose an application? 
Suggested Answer: B 🗳️ 

A. Port

B. Service

C. DNS

D. Deployment

**Answer: B**

**Timestamp: March 7, 2024, 5 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/135370-exam-kcna-topic-1-question-55-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B A Service is the Kubernetes resource used to expose an application running on a set of Pods. It provides a stable endpoint, load balancing, and service discovery, making it an essential component for accessing applications within a Kubernetes cluster. upvoted 1 times ... pulsefire 1 year, 4 months ago Selected Answer: B https://kubernetes.io/docs/tutorials/kubernetes-basics/expose/expose-intro/#:~:text=Expose%20Your%20App-,Using%20a%20Service%20to%20Expose%20Your%20App,-Objectives upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 56 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 56
Topic #: 1

[All KCNA Questions]

What is a DaemonSet? 
Suggested Answer: A 🗳️ 

A. It's a type of workload that ensures a specific set of nodes run a copy of a Pod.

B. It's a type of workload responsible for maintaining a stable set of replica Pods running in any node.

C. It's a type of workload that needs to be run periodically on a given schedule.

D. It's a type of workload that provides guarantees about ordering, uniqueness, and identity of a set of Pods.

**Answer: A**

**Timestamp: Sept. 28, 2024, 1 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/148291-exam-kcna-topic-1-question-56-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A A DaemonSet in Kubernetes is a type of workload that ensures a copy of a Pod runs on a specific set of nodes, typically all nodes in the cluster. DaemonSets are used for deploying system-level services that need to run on every node, such as log collectors, monitoring agents, or network plugins. upvoted 1 times ... 2211094 12 months ago Selected Answer: A A. It's a type of workload that ensures a specific set of nodes run a copy of a Pod. A DaemonSet ensures that all (or some) nodes run a copy of a specific Pod. When a new node is added to the cluster, a Pod from the DaemonSet is automatically added to that node. When a node is removed from the cluster, the Pod is also removed. upvoted 1 times ... mc2301 1 year, 3 months ago DaemonSet ensures that a copy of a Pod runs on all (or some) nodes of your cluster. DaemonSets are perfect to run infrastructure-related workload, for example monitoring or logging tools. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 57 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 57
Topic #: 1

[All KCNA Questions]

What is the telemetry component that represents a series of related distributed events that encode the end-to-end request flow through a distributed system? 
Suggested Answer: D 🗳️ 

A. Metrics

B. Logs

C. Spans

D. Traces

**Answer: D**

**Timestamp: Sept. 28, 2024, 1:04 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/148292-exam-kcna-topic-1-question-57-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D Traces are the telemetry component that represents a series of related distributed events, encoding the end-to-end request flow through a distributed system. They provide a comprehensive view of how requests propagate through various services and components, enabling detailed performance monitoring and troubleshooting. upvoted 1 times ... 2211094 12 months ago Selected Answer: D D. Traces In a distributed system, traces represent a series of related events that show the end-to-end request flow. Each trace is composed of one or more spans, which represent individual operations within the trace. upvoted 1 times ... mc2301 1 year, 3 months ago Traces track the progression of a request while it’s passing through the system. Traces are used in a distributed system that can provide information about when a request was processed by a service and how long it took. "Distributed" is the key here. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 58 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 58
Topic #: 1

[All KCNA Questions]

In the Kubernetes platform, which component is responsible for running containers? 
Suggested Answer: B 🗳️ 

A. etcd

B. CRI-O

C. cloud-controller-manager

D. kube-controller-manager

**Answer: B**

**Timestamp: Sept. 28, 2024, 1:11 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/148293-exam-kcna-topic-1-question-58-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B CRI-O is the component responsible for running containers in a Kubernetes cluster. It is a lightweight container runtime that implements the Kubernetes Container Runtime Interface (CRI), allowing it to manage the lifecycle of containers efficiently. upvoted 1 times ... 2211094 12 months ago Selected Answer: B B. CRI-O CRI-O is a lightweight container runtime specifically designed for Kubernetes. It implements the Kubernetes Container Runtime Interface (CRI) to enable the use of containers within Kubernetes clusters. It is responsible for running the containers specified by Kubernetes upvoted 1 times ... mc2301 1 year, 3 months ago Container Runtimes containerd is a lightweight and performant implementation to run containers. Arguably the most popular container runtime right now. It is used by all major cloud providers for the Kubernetes As A Service products. CRI-O was created by Red Hat and with a similar code base closely related to podman and buildah. Docker - The standard for a long time, but never really made for container orchestration. The usage of Docker as the runtime for Kubernetes has been deprecated and removed in Kubernetes 1.24. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 59 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 59
Topic #: 1

[All KCNA Questions]

Services and Pods in Kubernetes are ______ objects. 
Suggested Answer: D 🗳️ 

A. JSON

B. YAML

C. Java

D. REST

**Answer: D**

**Timestamp: March 7, 2024, 5:11 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/135371-exam-kcna-topic-1-question-59-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D Services and Pods in Kubernetes are API objects. They are defined and managed through the Kubernetes API, allowing users to declaratively configure and control their behavior within the cluster. upvoted 1 times ... 2211094 12 months ago Selected Answer: D B. CRI-O CRI-O is a lightweight container runtime specifically designed for Kubernetes. It implements the Kubernetes Container Runtime Interface (CRI) to enable the use of containers within Kubernetes clusters. It is responsible for running the containers specified by Kubernetes upvoted 1 times ... mc2301 1 year, 3 months ago The Kubernetes API is the most important component of a Kubernetes cluster. Without it, communication with the cluster is not possible, every user and every component of the cluster itself needs the api-server. Like many other APIs, the Kubernetes API is implemented as a RESTful interface that is exposed over HTTPS. Through the API, a user or service can create, modify, delete or retrieve resources that reside in Kubernetes. D seems legit upvoted 2 times ... 377ecc2 1 year, 7 months ago Selected Answer: D Services and Pods in Kubernetes are: D. REST objects In Kubernetes, resources such as Services and Pods are represented as RESTful objects. This means that they can be accessed and manipulated using HTTP methods such as GET, POST, PUT, and DELETE via the Kubernetes API server. While configurations for Kubernetes resources are often written in YAML or JSON format, these formats are used to define the desired state of the objects rather than representing the nature of the objects themselves upvoted 3 times ... JBangura 1 year, 9 months ago Selected Answer: D It's D upvoted 2 times ... pulsefire 1 year, 10 months ago Selected Answer: D should be D. REST objects https://docs.openshift.com/container-platform/3.11/architecture/core_concepts/pods_and_services.html#:~:text=Like%20pods%2C%20services%20are%20REST%20objects. upvoted 4 times ...
----------------------------------------

## Exam KCNA topic 1 question 60 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 60
Topic #: 1

[All KCNA Questions]

What Kubernetes component handles network communications inside and outside of a cluster, using operating system packet filtering if available? 
Suggested Answer: A 🗳️ 

A. kube-proxy

B. kubelet

C. etcd

D. kube-controller-manager

**Answer: A**

**Timestamp: March 7, 2024, 5:13 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/135372-exam-kcna-topic-1-question-60-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A kube-proxy is the Kubernetes component responsible for handling network communications inside and outside of a cluster. It maintains network rules on each node, which allow network traffic to be properly routed to and from the Pods. kube-proxy can use operating system packet filtering (such as iptables or IPVS) to manage these network rules. upvoted 1 times ... cajif66766 1 year ago Selected Answer: A A. kube-proxy upvoted 2 times ... Bot123 1 year, 5 months ago a is correct upvoted 1 times ... pulsefire 1 year, 10 months ago Selected Answer: A A https://kubernetes.io/docs/concepts/overview/components/#:~:text=kube%2Dproxy%20maintains,and%20it%27s%20available. upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 61 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 61
Topic #: 1

[All KCNA Questions]

What Kubernetes control plane component exposes the programmatic interface used to create, manage and interact with the Kubernetes objects? 
Suggested Answer: C 🗳️ 

A. kube-controller-manager

B. kube-proxy

C. kube-apiserver

D. etcd

**Answer: C**

**Timestamp: June 28, 2024, 12:52 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/143005-exam-kcna-topic-1-question-61-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C The kube-apiserver is the Kubernetes control plane component that exposes the programmatic interface used to create, manage, and interact with Kubernetes objects. It serves as the central management entity, processing API requests, handling authentication and authorization, and managing the cluster's state. upvoted 1 times ... 2211094 12 months ago Selected Answer: C C. kube-apiserver The kube-apiserver is the Kubernetes control plane component that exposes the Kubernetes API. It provides the programmatic interface used to create, manage, and interact with Kubernetes objects. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 62 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 62
Topic #: 1

[All KCNA Questions]

Which type of Service requires manual creation of Endpoints? 
Suggested Answer: B 🗳️ 

A. LoadBalancer

B. Services without selectors

C. NodePort

D. ClusterIP with selectors

**Answer: B**

**Timestamp: Jan. 12, 2025, 1:29 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/154427-exam-kcna-topic-1-question-62-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B Services without selectors require manual creation of Endpoints to specify the backend IP addresses and ports. This allows you to define custom backend services or external resources that the Service should route traffic to. upvoted 1 times ... 2211094 12 months ago Selected Answer: B B. Services without selectors When you create a Kubernetes Service without specifying selectors, Kubernetes does not automatically create Endpoints for the Service. In such cases, you need to manually create the Endpoints to map the Service to the set of Pods or IP addresses you want it to target upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 63 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 63
Topic #: 1

[All KCNA Questions]

Which of these commands is used to retrieve the documentation and field definitions for a Kubernetes resource? 
Suggested Answer: A 🗳️ 

A. kubectl explain

B. kubectl api-resources

C. kubectl get --help

D. kubectl show

**Answer: A**

**Timestamp: April 23, 2024, 9:30 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/139400-exam-kcna-topic-1-question-63-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A The kubectl explain command is used to retrieve the documentation and field definitions for a Kubernetes resource. It provides detailed information about the structure and usage of resources, helping users understand how to configure and interact with them. upvoted 1 times ... 2211094 12 months ago Selected Answer: A A. kubectl explain The kubectl explain command is used to retrieve documentation and field definitions for a Kubernetes resource. It provides detailed information about the fields and their descriptions for a specific resource type. upvoted 1 times ... hovnival 1 year, 2 months ago Selected Answer: A ChatGPT 3.5 says A "kubectl explain <resource-name>" Copilot says B even after some talk to it. I am going for A. upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 64 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 64
Topic #: 1

[All KCNA Questions]

Which of the following is a lightweight tool that manages traffic flows between services, enforces access policies, and aggregates telemetry data, all without requiring changes to application code? 
Suggested Answer: B 🗳️ 

A. NetworkPolicy

B. Linkerd

C. kube-proxy

D. Nginx

**Answer: B**

**Timestamp: Dec. 8, 2024, 4:20 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/152691-exam-kcna-topic-1-question-64-discussion/)

Comments: iamtrb 1 year, 1 month ago Selected Answer: B ServiceMesh, LinkerD has this kind of mechanism upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 65 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 65
Topic #: 1

[All KCNA Questions]

Which Kubernetes resource uses immutable: true boolean field? 
Suggested Answer: C 🗳️ 

A. Deployment

B. Pod

C. ConfigMap

D. ReplicaSet

**Answer: C**

**Timestamp: Dec. 6, 2023, 7:19 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/127939-exam-kcna-topic-1-question-65-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C The ConfigMap resource in Kubernetes can use the immutable: true boolean field to make the ConfigMap immutable, preventing any changes to its data after creation. This feature helps improve performance and ensures the stability of configuration data. upvoted 1 times ... 2211094 12 months ago Selected Answer: C C. ConfigMap In Kubernetes, you can create ConfigMaps with an immutable: true boolean field. This feature ensures that the data within the ConfigMap cannot be modified after creation, which can help improve the cluster's security and stability. upvoted 1 times ... alex78 1 year, 2 months ago Selected Answer: C https://kubernetes.io/docs/concepts/configuration/configmap/#configmap-immutable upvoted 3 times ... hovnival 1 year, 2 months ago Selected Answer: C Copilot says C giving also similiar example like Bolgarwow gave. upvoted 2 times ... Debayan 1 year, 3 months ago Selected Answer: C Kubernetes ConfigMap and Secret resources have a field called immutable which, if set to true, ensures that the data of ConfigMap/Secret cannot be updated after the resource is created. upvoted 3 times ... Bolgarwow 1 year, 3 months ago apiVersion: v1 kind: ConfigMap metadata: name: test-immutable labels: app.kubernetes.io/name: kubernetes-hacks-and-tricks app.kubernetes.io/created-by: ssbostan immutable: true ConfigMap upvoted 3 times ... omerco61 1 year, 5 months ago Selected Answer: A In Kubernetes, the immutable: true boolean field is associated with the Deployment resource. This field is used to indicate whether the Deployment is immutable, meaning that its spec, such as the template for the Pods it manages, cannot be changed. upvoted 1 times ... ankre 1 year, 7 months ago B is correct upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 66 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 66
Topic #: 1

[All KCNA Questions]

Which statement about the Kubernetes network model is correct? 
Suggested Answer: B 🗳️ 

A. Pods can only communicate with Pods exposed via a Service.

B. Pods can communicate with all Pods without NAT.

C. The Pod IP is only visible inside a Pod.

D. The Service IP is used for the communication between Services.

**Answer: B**

**Timestamp: March 7, 2024, 5:29 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/135374-exam-kcna-topic-1-question-66-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B The correct statement about the Kubernetes network model is that Pods can communicate with all Pods without NAT. This principle ensures a flat network topology where direct communication between Pods is possible, simplifying network configuration and management within the Kubernetes cluster. upvoted 1 times ... 2211094 12 months ago Selected Answer: B B. Pods can communicate with all Pods without NAT. In the Kubernetes network model, each Pod has a unique IP address, and Pods can communicate with each other directly without needing Network Address Translation (NAT). upvoted 1 times ... pulsefire 1 year, 4 months ago Selected Answer: B B https://kubernetes.io/docs/concepts/services-networking/#:~:text=pods%20can%20communicate%20with%20all%20other%20pods%20on%20any%20other%20node%20without%20NAT upvoted 3 times ...
----------------------------------------

## Exam KCNA topic 1 question 67 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 67
Topic #: 1

[All KCNA Questions]

What is the resource type used to package sets of containers for scheduling in a cluster? 
Suggested Answer: A 🗳️ 

A. Pod

B. ContainerSet

C. ReplicaSet

D. Deployment

**Answer: A**

**Timestamp: Jan. 12, 2025, 1:38 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/154428-exam-kcna-topic-1-question-67-discussion/)

Comments: 2211094 12 months ago Selected Answer: A Pod is the answer upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 68 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 68
Topic #: 1

[All KCNA Questions]

Can a Kubernetes Service expose multiple ports? 
Suggested Answer: B 🗳️ 

A. No, you can only expose one port per each Service.

B. Yes, but you must specify an unambiguous name for each port.

C. Yes, the only requirement is to use different port numbers.

D. No, because the only port you can expose is port number 443.

**Answer: B**

**Timestamp: Dec. 28, 2023, 10:23 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/129528-exam-kcna-topic-1-question-68-discussion/)

Comments: iulian0585 6 months, 2 weeks ago Selected Answer: C The name field for each port in the spec.ports array of a Kubernetes Service is optional. The Service will function correctly as long as the port numbers are unique and properly configured. However, including port names is a best practice for clarity and compatibility with other Kubernetes features upvoted 1 times ... donathon 7 months, 3 weeks ago Selected Answer: B A Kubernetes Service can expose multiple ports, and this is commonly used when: An application exposes multiple protocols (e.g., HTTP and gRPC) Different components need to be accessed on different ports (e.g., HTTP on 80, metrics on 8080) However, if multiple ports are defined, each port must have a unique name — especially when using protocols like HTTP/2 or gRPC, which rely on port names to route correctly (e.g., in service meshes like Istio or Linkerd). A. No, you can only expose one port per each Service: ❌ Incorrect — Services can expose multiple ports. C. Yes, the only requirement is to use different port numbers: ❌ Partially correct — different ports are required, but naming is also required for disambiguation. D. No, because the only port you can expose is port number 443: ❌ False — Services can expose any valid port number, not just 443. upvoted 2 times ... Matsuura 9 months, 2 weeks ago Selected Answer: C This discussion is very divided. But in a multiport definition, the name is optional. By networking logic, binding a port two times is not allowed. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: B a Service can expose multiple ports. When defining a Service with multiple ports, you must specify a unique name for each port to avoid ambiguity. This allows the Service to route traffic to different ports on the backend Pods. upvoted 1 times ... 2211094 12 months ago Selected Answer: B B. Yes, but you must specify an unambiguous name for each port. In Kubernetes, a Service can expose multiple ports, but you need to assign a unique name to each port to avoid ambiguity. upvoted 1 times ... abitwrong 1 year, 1 month ago Selected Answer: B https://kubernetes.io/docs/concepts/services-networking/service/#multi-port-services upvoted 2 times ... mc2301 1 year, 3 months ago Multi-port Services For some Services, you need to expose more than one port. Kubernetes lets you configure multiple port definitions on a Service object. When using multiple ports for a Service, you must give all of your ports names so that these are unambiguous. For example: apiVersion: v1 kind: Service metadata: name: my-service spec: selector: app.kubernetes.io/name: MyApp ports: - name: http protocol: TCP port: 80 targetPort: 9376 - name: https protocol: TCP port: 443 targetPort: 9377 upvoted 3 times ... miskill 1 year, 3 months ago Selected Answer: B B is correct: C suggests that only different port numbers are required, but it overlooks the need for a unique name to avoid confusion and conflicts, especially when multiple ports are involved. Therefore, B is correct because, in addition to different port numbers, a unique name for each port is required. upvoted 3 times ... miskill 1 year, 4 months ago Selected Answer: C ist is correct upvoted 1 times ... hovnival 1 year, 8 months ago Selected Answer: C Copilot says C upvoted 2 times ... pulsefire 1 year, 10 months ago Selected Answer: C should be C upvoted 2 times ... omerco61 1 year, 11 months ago Option B is incorrect because, while you can name the ports for clarity, it is not a strict requirement. Correct C upvoted 2 times ... 3c9bc24 2 years ago Isn't the correct answer C? I don't think it requires a name for each port upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 69 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 69
Topic #: 1

[All KCNA Questions]

Which persona is normally responsible for defining, testing, and running an incident management process? 
Suggested Answer: A 🗳️ 

A. Site Reliability Engineers

B. Project Managers

C. Application Developers

D. Quality Engineers

**Answer: A**

**Timestamp: Jan. 12, 2025, 1:44 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/154429-exam-kcna-topic-1-question-69-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A Site Reliability Engineers (SREs) are normally responsible for defining, testing, and running an incident management process upvoted 1 times ... 2211094 12 months ago Selected Answer: A A. Site Reliability Engineers Site Reliability Engineers (SREs) are typically responsible for defining, testing, and running incident management processes. They focus on ensuring the reliability and performance of systems, as well as handling operational issues and incidents. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 70 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 70
Topic #: 1

[All KCNA Questions]

What is the default deployment strategy in Kubernetes? 
Suggested Answer: A 🗳️ 

A. Rolling update

B. Blue/Green deployment

C. Canary deployment

D. Recreate deployment

**Answer: A**

**Timestamp: March 1, 2025, 4:16 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157295-exam-kcna-topic-1-question-70-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: A The default deployment strategy in Kubernetes is the Rolling Update. This strategy ensures that when an application is updated, new pods are gradually introduced while old pods are simultaneously terminated. This process allows for updates with minimal to no downtime, as a portion of the application remains available to serve traffic throughout the update process. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: A The default deployment strategy in Kubernetes is the rolling update strategy. This strategy ensures that updates are applied incrementally, maintaining application availability and minimizing downtime during the update process. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 71 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 71
Topic #: 1

[All KCNA Questions]

Which command provides information about the field replicas within the spec resource of a deployment object? 
Suggested Answer: B 🗳️ 

A. kubectl get deployment.spec.replicas

B. kubectl explain deployment.spec.replicas

C. kubectl describe deployment.spec.replicas

D. kubectl explain deployment --spec.replicas

**Answer: B**

**Timestamp: May 3, 2024, 10:23 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/139950-exam-kcna-topic-1-question-71-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B kubectl explain deployment.spec.replicas upvoted 1 times ... miskill 1 year, 3 months ago Selected Answer: B The correct answer is: B. kubectl explain deployment.spec.replicas Explanation: The command kubectl explain is used to provide detailed information about the fields of Kubernetes objects, including their specifications. The command kubectl explain deployment.spec.replicas specifically provides information about the replicas field within the spec of a deployment object. This is the most appropriate command for getting details about that particular field. upvoted 1 times ... EzBL 1 year, 6 months ago Selected Answer: B A. kubectl get deployment.spec.replicas: This command attempts to retrieve the value of replicas directly from the deployment resource. However, it's not a standard way to access nested fields like spec.replicas. B. kubectl explain deployment.spec.replicas: This command is used to get documentation or information about a specific field within a Kubernetes resource definition. It will provide details about what the replicas field represents within the spec of a deployment. C. kubectl describe deployment.spec.replicas: This command describes an existing resource, such as a deployment, but it does not specifically explain the meaning or purpose of a field like replicas within the spec. D. kubectl explain deployment --spec.replicas: This command is close but not entirely correct. It attempts to explain a flag --spec.replicas for the kubectl explain command, rather than directly explaining the replicas field within the spec of a deployment. upvoted 2 times ... alex78 1 year, 8 months ago Selected Answer: B https://kubernetes.io/docs/reference/kubectl/generated/kubectl_explain/ upvoted 2 times ... hovnival 1 year, 8 months ago Selected Answer: B guys this one is totally B. kubectl explain deployment.spec.replicas kubectl explain deployment => info about deployment kubectl explain deployment.spec => info about spec in deployment kubectl explain deployment.spec.replicas => info aout replicas in spec in deployment upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 72 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 72
Topic #: 1

[All KCNA Questions]

Which of the following is a responsibility of the governance board of an open source project? 
Suggested Answer: C 🗳️ 

A. Decide about the marketing strategy of the project.

B. Review the pull requests in the main branch.

C. Outline the project's “terms of engagement”.

D. Define the license to be used in the project.

**Answer: C**

**Timestamp: June 2, 2024, 6:37 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/141758-exam-kcna-topic-1-question-72-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C The governance board of an open source project is responsible for outlining the project's "terms of engagement." This includes defining the rules, guidelines, and processes for how contributors and maintainers interact with the project. The terms of engagement help ensure that the project operates smoothly, transparently, and inclusively. upvoted 1 times ... 2211094 12 months ago Selected Answer: C C. Outline the project's “terms of engagement”. The governance board of an open source project is responsible for setting the project's "terms of engagement," which includes defining the rules, guidelines, and processes for contributing to and interacting with the project. This ensures that the project operates smoothly and transparently. upvoted 1 times ... abu7midan 1 year, 1 month ago Governing Board is responsible for marketing and other business oversight and budget decisions for the CNCF i think answer is A upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 73 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 73
Topic #: 1

[All KCNA Questions]

What is the role of a NetworkPolicy in Kubernetes? 
Suggested Answer: B 🗳️ 

A. The ability to cryptic and obscure all traffic.

B. The ability to classify the Pods as isolated and non isolated.

C. The ability to prevent loopback or incoming host traffic.

D. The ability to log network security events.

**Answer: B**

**Timestamp: Feb. 10, 2024, 1:53 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/133487-exam-kcna-topic-1-question-73-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B A NetworkPolicy in Kubernetes is a resource that defines rules for controlling the network traffic to and from Pods. NetworkPolicies allow you to specify how groups of Pods are allowed to communicate with each other and with other network endpoints. By default, Pods are non-isolated, meaning they accept traffic from any source. When a NetworkPolicy is applied, it can classify Pods as isolated, meaning they only accept traffic that matches the specified rules. upvoted 1 times ... 2211094 12 months ago Selected Answer: B B. The ability to classify the Pods as isolated and non isolated. In Kubernetes, a NetworkPolicy is used to control the communication between Pods within the cluster. It allows you to specify rules for which Pods can communicate with each other, effectively classifying Pods as isolated or non-isolated based on the defined rules upvoted 1 times ... pulsefire 1 year, 4 months ago Selected Answer: B B https://kubernetes.io/docs/concepts/services-networking/network-policies/#the-two-sorts-of-pod-isolation:~:text=There%20are%20two,pod%20to%20another. upvoted 2 times ... omerco61 1 year, 5 months ago Selected Answer: B B corret. https://kubernetes.io/docs/concepts/services-networking/network-policies/#the-two-sorts-of-pod-isolation upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 74 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 74
Topic #: 1

[All KCNA Questions]

What are the most important resources to guarantee the performance of an etcd cluster? 
Suggested Answer: B 🗳️ 

A. CPU and disk capacity.

B. Network throughput and disk I/O.

C. CPU and RAM memory.

D. Network throughput and CPU.

**Answer: B**

**Timestamp: Dec. 6, 2023, 7:41 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/127942-exam-kcna-topic-1-question-74-discussion/)

Comments: ankre Highly Voted 2 years, 1 month ago B is correct upvoted 7 times ... shahy0 Most Recent 10 months, 2 weeks ago Selected Answer: B The most important resources to guarantee the performance of an etcd cluster are network throughput and disk I/O. High network throughput ensures fast and reliable communication between etcd nodes, while high disk I/O performance ensures efficient write and read operations, both of which are crucial for maintaining the responsiveness and reliability of the etcd cluster. upvoted 1 times ... BvGVAXeAMP 1 year, 4 months ago B is the correct answer upvoted 2 times ... pulsefire 1 year, 10 months ago Selected Answer: B should be B upvoted 3 times ... omerco61 1 year, 11 months ago Consensus performance, especially commit latency, is limited by two physical constraints: network IO latency and disk IO latency. upvoted 3 times ... mlsafonseca 2 years ago It's B! upvoted 2 times ... AbhishekJoshi 2 years ago Selected Answer: B https://etcd.io/docs/v3.4/op-guide/performance/ upvoted 4 times ...
----------------------------------------

## Exam KCNA topic 1 question 75 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 75
Topic #: 1

[All KCNA Questions]

How do you deploy a workload to Kubernetes without additional tools? 
Suggested Answer: C 🗳️ 

A. Create a Bash script and run it on a worker node.

B. Create a Helm Chart and install it with helm.

C. Create a manifest and apply it with kubectl.

D. Create a Python script and run it with kubectl.

**Answer: C**

**Timestamp: March 1, 2025, 4:48 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157296-exam-kcna-topic-1-question-75-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C To deploy a workload to Kubernetes without additional tools, you can create a manifest file (typically in YAML format) that defines the desired state of your Kubernetes resources (such as Pods, Deployments, Services, etc.) and then apply this manifest using the kubectl command-line tool. upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 76 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 76
Topic #: 1

[All KCNA Questions]

How do you perform a command in a running container of a Pod? 
Suggested Answer: A 🗳️ 

A. kubectl exec --

B. docker exec

C. kubectl run --

D. kubectl attach -i

**Answer: A**

**Timestamp: March 20, 2024, 10:43 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/136713-exam-kcna-topic-1-question-76-discussion/)

Comments: themickou 5 months ago Selected Answer: A kubectl exec -it nginx -- bash upvoted 1 times ... cajif66766 1 year ago Selected Answer: A kubectl pod_name exec -- upvoted 1 times ... seetpt 1 year, 4 months ago Selected Answer: A A is correct upvoted 1 times ... wdwe 1 year, 9 months ago Should be B upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 77 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 77
Topic #: 1

[All KCNA Questions]

How to create a headless service? 
Suggested Answer: B 🗳️ 

A. By specifying .spec.ClusterIP: headless

B. By specifying .spec.clusterIP: None

C. By specifying .spec.ClusterIP: 0.0.0.0

D. By specifying .spec.ClusterIP: localhost

**Answer: B**

**Timestamp: Sept. 18, 2024, 10:51 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/147802-exam-kcna-topic-1-question-77-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B To create a headless Service in Kubernetes, you need to specify .spec.clusterIP: None in the Service manifest. This configuration allows clients to directly access the individual Pods' IP addresses without using a ClusterIP. upvoted 1 times ... vspringe 1 year, 3 months ago Selected Answer: B B. By specifying .spec.clusterIP: None upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 78 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 78
Topic #: 1

[All KCNA Questions]

How does dynamic storage provisioning work? 
Suggested Answer: A 🗳️ 

A. A user requests dynamically provisioned storage by including an existing storage class in their PersistentVolumeClaim.

B. An administrator creates a storage class and includes it in their pod YAML definition file without creating a PersistentVolumeClaim.

C. A pod requests dynamically provisioned storage by including a storage class and the pod name in their PersistentVolumeClaim.

D. An administrator creates a PersistentVolume and includes the name of the PersistentVolume in their pod YAML definition file.

**Answer: A**

**Timestamp: March 7, 2024, 5:56 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/135377-exam-kcna-topic-1-question-78-discussion/)

Comments: pulsefire 1 year, 4 months ago Selected Answer: A A https://kubernetes.io/docs/concepts/storage/dynamic-provisioning/#:~:text=The%20dynamic%20provisioning%20feature%20eliminates,it%20is%20requested%20by%20users. upvoted 3 times ...
----------------------------------------

## Exam KCNA topic 1 question 79 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 79
Topic #: 1

[All KCNA Questions]

Which of the following are tasks performed by a container orchestration tool? 
Suggested Answer: A 🗳️ 

A. Schedule, scale, and manage the health of containers.

B. Create images, scale, and manage the health of containers.

C. Debug applications, and manage the health of containers.

D. Store images, scale, and manage the health of containers.

**Answer: A**

**Timestamp: Dec. 9, 2024, 7:55 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/152759-exam-kcna-topic-1-question-79-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A Container orchestration tools, such as Kubernetes, are responsible for scheduling containers on nodes, scaling the number of container instances, and managing the health of containers. These tasks ensure that containerized applications run efficiently, reliably, and can handle varying workloads. upvoted 1 times ... mainevent55 1 year, 1 month ago Selected Answer: A A. https://cloud.google.com/discover/what-is-container-orchestration?hl=en upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 80 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 80
Topic #: 1

[All KCNA Questions]

Which of the following is a definition of Hybrid Cloud? 
Suggested Answer: C 🗳️ 

A. An architecture that uses a combination of services running in public and private data centers, only including data centers from the same cloud provider.

B. A cloud native architecture that uses a combination of services running in public clouds, excluding data centers in different availability zones.

C. A cloud native architecture that uses a combination of services running in different public and private clouds, including on-premises data centers.

D. An architecture that uses a combination of services running in public and private data centers, excluding serverless functions.

**Answer: C**

**Timestamp: Jan. 12, 2025, 2:11 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/154432-exam-kcna-topic-1-question-80-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C A Hybrid Cloud is a cloud native architecture that uses a combination of services running in different public and private clouds, including on-premises data centers. This approach provides flexibility, scalability, and enhanced security by leveraging the strengths of both public and private cloud environments. upvoted 1 times ... 2211094 12 months ago Selected Answer: C C. A cloud native architecture that uses a combination of services running in different public and private clouds, including on-premises data centers. A Hybrid Cloud architecture combines resources from public clouds, private clouds, and on-premises data centers, allowing for greater flexibility and scalability while leveraging the benefits of both environments. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 81 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 81
Topic #: 1

[All KCNA Questions]

What is a Kubernetes Service Endpoint? 
Suggested Answer: D 🗳️ 

A. It is the API Endpoint of our Kubernetes cluster.

B. It is a name of special Pod in kube-system namespace.

C. It is an IP address that we can access from the Internet.

D. It is an object that gets IP addresses of individual Pods assigned to it.

**Answer: D**

**Timestamp: March 7, 2024, 3:52 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/135367-exam-kcna-topic-1-question-81-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D A Kubernetes Service Endpoint is a resource that provides the actual IP addresses and ports of the Pods that are backing a Service. Endpoints are dynamically created and updated by Kubernetes to reflect the current state of the Pods that match the Service's selector. upvoted 1 times ... 2211094 12 months ago Selected Answer: D D is the correct answer. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 82 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 82
Topic #: 1

[All KCNA Questions]

Why is Cloud-Native Architecture important? 
Suggested Answer: B 🗳️ 

A. Cloud Native Architecture revolves around containers, microservices and pipelines.

B. Cloud Native Architecture removes constraints to rapid innovation.

C. Cloud Native Architecture is modern for application deployment and pipelines.

D. Cloud Native Architecture is a bleeding edge technology and service.

**Answer: B**

**Timestamp: April 25, 2024, 9:13 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/139560-exam-kcna-topic-1-question-82-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B Cloud-Native Architecture is important because it removes constraints to rapid innovation. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: C Cloud-Native Architecture is important for several reasons, particularly in the context of modern application development and deployment upvoted 1 times ... hovnival 1 year, 2 months ago Selected Answer: B Copilot says A, ChatGPT says B. I am going for B as it sounds like promo for Cloud-Native Architecture, right? upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 83 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 83
Topic #: 1

[All KCNA Questions]

Which kubernetes component is the smallest deployable unit of computing? 
Suggested Answer: C 🗳️ 

A. StatefulSet

B. Deployment

C. Pod

D. Container

**Answer: C**

**Timestamp: Dec. 7, 2023, 5:25 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/128035-exam-kcna-topic-1-question-83-discussion/)

Comments: ankre Highly Voted 2 years, 1 month ago C is correct, container is not deployable upvoted 5 times ... shahy0 Most Recent 10 months, 2 weeks ago Selected Answer: C A Pod is the smallest and simplest unit in the Kubernetes object model that you can create or deploy upvoted 1 times ... miskill 1 year, 3 months ago Selected Answer: C The correct answer is: C. Pod Explanation: In Kubernetes, the Pod is the smallest deployable unit of computing. A Pod can contain one or more containers, but it is the fundamental unit that Kubernetes deploys, manages, and scales. Containers within a Pod share the same network and storage resources, making the Pod the atomic unit for scheduling and running containers in Kubernetes. upvoted 1 times ... SeaH0rse66 1 year, 7 months ago Selected Answer: C A Pod is the smallest deployable unit of computing in Kubernetes. It represents a single instance of a running process in your cluster. Pods can contain one or more containers that are tightly coupled and share resources, such as network and storage, and are scheduled onto nodes in the Kubernetes cluster., this is in a way a "unit of measure" , like the atom of Kubernetes, the container would be like a proton... upvoted 2 times ... andytsangchun 1 year, 8 months ago For sure its C upvoted 3 times ... hovnival 1 year, 8 months ago Selected Answer: C C is correct, confirmed by Copilot. upvoted 2 times ... EzBL 1 year, 9 months ago Selected Answer: C Explanation: • In Kubernetes, the smallest deployable unit of computing is a Pod. • A Pod represents a single instance of a running process in your cluster. • It can contain one or more containers that are tightly coupled and share the same network namespace, IPC namespace, and other resources. • Pods are the basic building blocks of Kubernetes applications and are scheduled onto nodes in the cluster. upvoted 3 times ... pulsefire 1 year, 10 months ago Selected Answer: C c is right upvoted 4 times ... majkisermi98 1 year, 12 months ago Selected Answer: C That's a basic knowledge upvoted 4 times ... 3c9bc24 2 years ago agree, C is correct upvoted 4 times ... AbhishekJoshi 2 years ago Selected Answer: C Container is not a object in k8s, pod is the only smallest unit in k8s upvoted 4 times ...
----------------------------------------

## Exam KCNA topic 1 question 85 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 85
Topic #: 1

[All KCNA Questions]

Which are the two primary modes for Service discovery within a Kubernetes cluster? 
Suggested Answer: A 🗳️ 

A. Environment variables and DNS

B. API Calls and LDAP

C. Labels and Radius

D. Selectors and DHCP

**Answer: A**

**Timestamp: Dec. 17, 2023, 8:53 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/128782-exam-kcna-topic-1-question-85-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A The two primary modes for Service discovery within a Kubernetes cluster are DNS-based Service Discovery and Environment Variable-based Service Discovery. DNS-based Service Discovery uses the built-in DNS server to provide DNS names for Services, while Environment Variable-based Service Discovery injects environment variables containing Service information into Pods. These mechanisms enable Pods to discover and communicate with Services within the cluster. upvoted 1 times ... 2211094 12 months ago Selected Answer: A Environment variables and DNS is the answer. upvoted 1 times ... AbhishekJoshi 1 year, 6 months ago Selected Answer: A https://www.densify.com/kubernetes-autoscaling/kubernetes-service-discovery/ .. please see conclusion upvoted 4 times ...
----------------------------------------

## Exam KCNA topic 1 question 86 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 86
Topic #: 1

[All KCNA Questions]

Which of the following capabilities are you allowed to add to a container using the Restricted policy? 
Suggested Answer: D 🗳️ 

A. CHOWN

B. SYS_CHROOT

C. SETUID

D. NET_BIND_SERVICE

**Answer: D**

**Timestamp: March 7, 2024, 4:15 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/135368-exam-kcna-topic-1-question-86-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D The Restricted policy in Kubernetes is designed to enforce strict security constraints on containers, but it does allow certain capabilities that are considered safe and necessary for common container operations. Among the capabilities listed, NET_BIND_SERVICE is typically allowed under the Restricted policy because it is often needed for binding to network ports below 1024, which is a common requirement for many applications upvoted 1 times ... d759fb3 10 months, 3 weeks ago Selected Answer: B The Restricted policy in Kubernetes is designed to be very secure, so it only allows a minimal set of capabilities. SYS_CHROOT is considered safe enough because: - It doesn't give elevated privileges - It's often needed for container operations - It can't be easily exploited upvoted 1 times ... yoyo2424 11 months, 3 weeks ago Selected Answer: A In Kubernetes, the Restricted policy is a security context that enforces tight constraints on what a container can do. When using the Restricted policy, the capabilities granted to containers are minimal and restricted to those necessary for basic functionality. The CHOWN capability is typically allowed in restricted environments because it is essential for many applications to modify file ownership within the container. upvoted 1 times ... 2211094 12 months ago Selected Answer: D D. NET_BIND_SERVICE The Restricted policy in Kubernetes is designed to limit the capabilities that can be added to containers to enhance security. Among the options provided, NET_BIND_SERVICE is the capability allowed under the Restricted policy. upvoted 1 times ... abitwrong 1 year, 1 month ago Selected Answer: D Containers must drop ALL capabilities, and are only permitted to add back the NET_BIND_SERVICE capability. Reference: https://kubernetes.io/docs/concepts/security/pod-security-standards/#:~:text=add%20back%20the-,NET_BIND_SERVICE,-capability.%20This upvoted 1 times ... Andrei_Z 1 year, 3 months ago Selected Answer: D https://kubernetes.io/docs/concepts/security/pod-security-standards/ Capabilities (v1.22+) Containers must drop ALL capabilities, and are only permitted to add back the NET_BIND_SERVICE capability. This is Linux only policy in v1.25+ (.spec.os.name != "windows") upvoted 2 times ... alex78 1 year, 8 months ago Selected Answer: D https://docs.openshift.com/dedicated/authentication/managing-security-context-constraints.html upvoted 3 times ... fabianvera19822 1 year, 8 months ago Option : A upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 87 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 87
Topic #: 1

[All KCNA Questions]

What methods can you use to scale a deployment? 
Suggested Answer: C 🗳️ 

A. With kubectl edit deployment exclusively.

B. With kubectl scale-up deployment exclusively.

C. With kubectl scale deployment and kubectl edit deployment.

D. With kubectl scale deployment exclusively.

**Answer: C**

**Timestamp: Jan. 12, 2025, 9:22 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/154441-exam-kcna-topic-1-question-87-discussion/)

Comments: 2211094 12 months ago Selected Answer: C C. With kubectl scale deployment and kubectl edit deployment. You can scale a Kubernetes deployment using either of the following methods: kubectl scale deployment: This command allows you to specify the desired number of replicas directly. kubectl edit deployment: This command opens the deployment's configuration in an editor, where you can manually change the number of replicas. Both methods effectively achieve the same result by adjusting the number of running instances (replicas) of your application. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 89 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 89
Topic #: 1

[All KCNA Questions]

Which is an industry-standard container runtime with an “emphasis” on simplicity, robustness, and portability? 
Suggested Answer: C 🗳️ 

A. cri-o

B. lxd

C. containerd

D. kata-runtime

**Answer: C**

**Timestamp: Dec. 9, 2024, 8:08 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/152761-exam-kcna-topic-1-question-89-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C containerd is an industry-standard container runtime that emphasizes simplicity, robustness, and portability. It is a core component of Docker and is also used by Kubernetes as a container runtime. containerd provides a simple and stable interface for managing container lifecycles, including image transfer, container execution, and storage upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 90 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 90
Topic #: 1

[All KCNA Questions]

What does vertical scaling an application deployment describe best? 
Suggested Answer: C 🗳️ 

A. The act of adding/removing applications to meet demand.

B. The act of adding/removing node instances to the cluster to meet demand.

C. The act of adding/removing resources to applications to meet demand.

D. The act of adding/removing application instances of the same application to meet demand.

**Answer: C**

**Timestamp: March 1, 2025, 3:13 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157324-exam-kcna-topic-1-question-90-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C Vertical scaling an application deployment describes the act of adding or removing resources (such as CPU, memory, or storage) to a single instance of an application to meet demand. This approach adjusts the capacity of a single application instance, in contrast to horizontal scaling, which involves adding or removing multiple instances of the application. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 92 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 92
Topic #: 1

[All KCNA Questions]

What is Serverless computing? 
Suggested Answer: A 🗳️ 

A. A computing method of providing backend services on an as-used basis.

B. A computing method of providing services for AI and ML operating systems.

C. A computing method of providing services for quantum computing operating systems.

D. A computing method of providing services for cloud computing operating systems.

**Answer: A**

**Timestamp: Jan. 12, 2025, 9:27 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/154442-exam-kcna-topic-1-question-92-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A Serverless computing is a cloud computing execution model where the cloud provider dynamically manages the allocation and provisioning of servers. In this model, developers can build and run applications without having to manage the underlying infrastructure. Instead, they focus on writing code, and the cloud provider handles the execution, scaling, and maintenance of the backend services on an as-used basis. upvoted 1 times ... 2211094 12 months ago Selected Answer: A A. A computing method of providing backend services on an as-used basis. Serverless computing is a cloud computing model in which the cloud provider automatically manages the infrastructure and dynamically allocates resources as needed. Developers can focus on writing code without worrying about server management, and they are only charged for the actual usage of resources rather than pre-provisioned capacity. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 93 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 93
Topic #: 1

[All KCNA Questions]

What is the purpose of the CRI? 
Suggested Answer: C 🗳️ 

A. To provide runtime integration control when multiple runtimes are used.

B. Support container replication and scaling on nodes.

C. Provide an interface allowing Kubernetes to support pluggable container runtimes.

D. Allow the definition of dynamic resource criteria across containers.

**Answer: C**

**Timestamp: March 1, 2025, 3:22 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157325-exam-kcna-topic-1-question-93-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C The Container Runtime Interface (CRI) is a standard API that allows Kubernetes to interact with different container runtimes in a consistent manner. The purpose of the CRI is to provide an abstraction layer that enables Kubernetes to support pluggable container runtimes, allowing users to choose the container runtime that best fits their needs. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 95 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 95
Topic #: 1

[All KCNA Questions]

In CNCF, who develops specifications for industry standards around container formats and runtimes? 
Suggested Answer: A 🗳️ 

A. Open Container Initiative (OCI)

B. Linux Foundation Certification Group (LFCG)

C. Container Network Interface (CNI)

D. Container Runtime Interface (CRI)

**Answer: A**

**Timestamp: March 1, 2025, 3:23 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157326-exam-kcna-topic-1-question-95-discussion/)

Comments: Matsuura 9 months, 2 weeks ago Selected Answer: A Open Container Initiative (OCI) born in the need of to have an open standard for container technologies which supports the collaboration in open source projects and prevent vendor lock-in. (https://opencontainers.org/) upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: A In the Cloud Native Computing Foundation (CNCF), the Open Container Initiative (OCI) is responsible for developing specifications for industry standards around container formats and runtimes upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 96 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 96
Topic #: 1

[All KCNA Questions]

Which of the following options includes valid API versions? 
Suggested Answer: C 🗳️ 

A. alpha1v1, beta3v3, v2

B. alpha1, beta3, v2

C. v1alpha1, v2beta3, v2

D. v1alpha1, v2beta3, 2.0

**Answer: C**

**Timestamp: Dec. 7, 2023, 6:06 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/128036-exam-kcna-topic-1-question-96-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C vXalphaY: Indicates an alpha version, where X is the major version and Y is the alpha version number. vXbetaY: Indicates a beta version, where X is the major version and Y is the beta version number. vX: Indicates a stable version, where X is the major version number. upvoted 1 times ... 2211094 12 months ago Selected Answer: C v1alpha1, v2beta3, v2 In Kubernetes, valid API versions follow the format v[MAJOR][MINOR][STABILITY], where: v1alpha1 represents an alpha version of the API, indicating an early stage and potentially unstable version. v2beta3 represents a beta version of the API, indicating a more stable version that is not yet finalized. v2 represents a stable version of the API. upvoted 1 times ... miskill 1 year, 3 months ago Selected Answer: C The correct answer is: C. v1alpha1, v2beta3, v2 Explanation: Kubernetes follows a specific convention for API versioning, which includes versions like v1alpha1, v2beta3, and stable versions like v1, v2, etc. The format typically consists of a version (vX), followed by an optional alpha or beta designation with a version number (e.g., alpha1, beta3). This structure is used to represent different stages of API stability, from alpha (experimental) to beta (more stable) to full release (e.g., v1, v2). upvoted 2 times ... SeaH0rse66 1 year, 7 months ago Selected Answer: C https://kubernetes.io/docs/reference/using-api/ Alpha: The version names contain alpha (for example, v1alpha1) Beta: The version names contain beta (for example, v2beta3). Stable: The version name is vX where X is an integer. upvoted 2 times ... andytsangchun 1 year, 8 months ago Selected Answer: C Its C From official docs: https://kubernetes.io/docs/reference/using-api/ upvoted 2 times ... hovnival 1 year, 8 months ago Selected Answer: C Copilot says C upvoted 2 times ... pulsefire 1 year, 10 months ago Selected Answer: C c is right upvoted 2 times ... pulsefire 1 year, 10 months ago Selected Answer: C c from chatgpt upvoted 2 times ... majkisermi98 1 year, 11 months ago Selected Answer: C C is the correct one upvoted 2 times ... AbhishekJoshi 2 years ago Selected Answer: A alpha1v1," "beta3v3," and "v2" follow the common pattern of API versioning with a "v" followed by a version number. upvoted 1 times ... ankre 2 years, 1 month ago C is correct. B. have an incorrect sequence of prefixes and version numbers. upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 97 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 97
Topic #: 1

[All KCNA Questions]

Which of the following will view the snapshot of previously terminated ruby container logs from Pod web-1? 
Suggested Answer: A 🗳️ 

A. kubectl logs -p -c ruby web-1

B. kubectl logs -c ruby web-1

C. kubectl logs -p ruby web-1

D. kubectl logs -p -c web-1 ruby

**Answer: A**

**Timestamp: Dec. 28, 2023, 7:44 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/129525-exam-kcna-topic-1-question-97-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A The kubectl logs command is used to retrieve logs from containers in Kubernetes Pods. To view the logs of a previously terminated container, you use the -p (previous) flag. To specify the container within the Pod, you use the -c (container) flag. upvoted 1 times ... 2211094 12 months ago Selected Answer: A kubectl logs -p -c ruby web-1 upvoted 1 times ... pulsefire 1 year, 4 months ago Selected Answer: A -p previous -c container upvoted 3 times ... dimoz 1 year, 6 months ago https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands#logs: -- Return snapshot of previous terminated ruby container logs from pod web-1 kubectl logs -p -c ruby web-1 upvoted 1 times ... ncsiki 1 year, 6 months ago Return snapshot of previous terminated ruby container logs from pod web-1 kubectl logs -p -c ruby web-1 https://jamesdefabia.github.io/docs/user-guide/kubectl/kubectl_logs/ upvoted 1 times ... PinkAndBlack 1 year, 6 months ago The correct syntax should be the following: kubectl logs web-1 -c ruby Check https://kubernetes.io/docs/reference/kubectl/quick-reference/#interacting-with-running-pods upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 98 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 98
Topic #: 1

[All KCNA Questions]

A Kubernetes _____ is an abstraction that defines a logical set of Pods and a policy by which to access them. 
Suggested Answer: C 🗳️ 

A. Selector

B. Controller

C. Service

D. Job

**Answer: C**

**Timestamp: Jan. 12, 2025, 9:42 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/154443-exam-kcna-topic-1-question-98-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C In Kubernetes, a Service is a resource that provides a stable endpoint (IP address and DNS name) to access a set of Pods. Services abstract away the details of individual Pods, allowing clients to interact with a consistent endpoint even as Pods are created, destroyed, or rescheduled. upvoted 1 times ... 2211094 12 months ago Selected Answer: C The answer is "Service" upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 99 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 99
Topic #: 1

[All KCNA Questions]

How many hosts are required to set up a highly available Kubernetes cluster when using an external etcd topology? 
Suggested Answer: D 🗳️ 

A. Four hosts. Two for control plane nodes and two for etcd nodes.

B. Four hosts. One for a control plane node and three for etcd nodes.

C. Three hosts. The control plane nodes and etcd nodes share the same host.

D. Six hosts. Three for control plane nodes and three for etcd nodes.

**Answer: D**

**Timestamp: Oct. 25, 2024, 10:33 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/150236-exam-kcna-topic-1-question-99-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D A highly available Kubernetes cluster with an external etcd topology requires separate hosts for the Kubernetes control plane components and the etcd cluster. Here is the breakdown: Kubernetes Control Plane Nodes: At least 3 hosts are recommended for the control plane components (kube-apiserver, kube-controller-manager, kube-scheduler) to ensure high availability. This allows for redundancy and failover in case one or more control plane nodes go down. etcd Cluster Nodes: At least 3 hosts are recommended for the etcd cluster to ensure high availability and data consistency. An odd number of etcd nodes is preferred to maintain quorum and avoid split-brain scenarios. upvoted 1 times ... 2211094 12 months ago Selected Answer: D D. Six hosts. Three for control plane nodes and three for etcd nodes. In a highly available Kubernetes cluster with an external etcd topology, you typically need three control plane nodes and three etcd nodes. This ensures redundancy and fault tolerance for both the control plane and the etcd database, providing high availability for the cluster. upvoted 1 times ... mainevent55 1 year, 1 month ago Selected Answer: D You should therefore run a minimum of three stacked control plane nodes for an HA cluster https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/ha-topology/#:~:text=A%20minimum%20of%20three%20hosts,HA%20cluster%20with%20this%20topology. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 100 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 100
Topic #: 1

[All KCNA Questions]

Which of these events will cause the kube-scheduler to assign a Pod to a node? 
Suggested Answer: D 🗳️ 

A. When the Pod crashes because of an error.

B. When a new node is added to the Kubernetes cluster.

C. When the CPU load on the node becomes too high.

D. When a new Pod is created and has no assigned node.

**Answer: D**

**Timestamp: March 1, 2025, 3:34 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157327-exam-kcna-topic-1-question-100-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D The kube-scheduler is responsible for assigning Pods to nodes in a Kubernetes cluster. It does this by watching for newly created Pods that do not have an assigned node and then selecting an appropriate node for them based on resource availability, constraints, and policies. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 102 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 102
Topic #: 1

[All KCNA Questions]

Which resource do you use to attach a volume in a Pod? 
Suggested Answer: D 🗳️ 

A. StorageVolume

B. PersistentVolume

C. StorageClass

D. PersistentVolumeClaim

**Answer: D**

**Timestamp: Jan. 12, 2025, 9:50 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/154444-exam-kcna-topic-1-question-102-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D In Kubernetes, a PersistentVolumeClaim (PVC) is used to request and attach a volume to a Pod. A PVC is a resource that allows users to request storage resources without needing to know the details of the underlying storage infrastructure. The PVC is then bound to a PersistentVolume (PV), which represents the actual storage resource. upvoted 1 times ... 2211094 12 months ago Selected Answer: D D. PersistentVolumeClaim In Kubernetes, a PersistentVolumeClaim (PVC) is used to request and attach a volume to a Pod. A PVC is a resource that allows users to dynamically provision storage based on a StorageClass or bind to an existing PersistentVolume. When a Pod specifies a PVC, the corresponding volume is attached to the Pod, providing persistent storage. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 103 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 103
Topic #: 1

[All KCNA Questions]

Which key-value store is used to persist Kubernetes cluster data? 
Suggested Answer: A 🗳️ 

A. etcd

B. ZooKeeper

C. ControlPlaneStore

D. Redis

**Answer: A**

**Timestamp: March 1, 2025, 3:38 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157329-exam-kcna-topic-1-question-103-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A etcd is the key-value store used to persist Kubernetes cluster data. It is a distributed, reliable, and consistent store that plays a critical role in maintaining the configuration and state information of the Kubernetes cluster. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 104 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 104
Topic #: 1

[All KCNA Questions]

What Linux namespace is shared by default by containers running within a Kubernetes Pod? 
Suggested Answer: B 🗳️ 

A. Host Network

B. Network

C. Process ID

D. Process Name

**Answer: B**

**Timestamp: Jan. 12, 2025, 9:52 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/154445-exam-kcna-topic-1-question-104-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B By default, containers running within the same Kubernetes Pod share the network namespace. This means that all containers in a Pod share the same IP address and port space, allowing them to communicate with each other using localhost and to coordinate on port usage. upvoted 1 times ... 2211094 12 months ago Selected Answer: B B. Network By default, containers running within a Kubernetes Pod share the same network namespace. This means that they can communicate with each other using localhost and can also share the same IP address and port space. Sharing the network namespace allows containers within the same Pod to efficiently collaborate and communicate. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 105 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 105
Topic #: 1

[All KCNA Questions]

What is a Dockerfile? 
Suggested Answer: C 🗳️ 

A. A bash script that is used to automatically build a docker image.

B. A config file that defines which image registry a container should be pushed to.

C. A text file that contains all the commands a user could call on the command line to assemble an image.

D. An image layer created by a running container stored on the host.

**Answer: C**

**Timestamp: March 1, 2025, 3:40 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157331-exam-kcna-topic-1-question-105-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C A Dockerfile is a text file that contains a set of instructions used to build a Docker image. It defines the steps needed to create a containerized application, including the base image, application code, dependencies, and configuration settings. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 106 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 106
Topic #: 1

[All KCNA Questions]

What does the "nodeSelector" within a PodSpec use to place Pods on the target nodes? 
Suggested Answer: D 🗳️ 

A. Annotations

B. IP Addresses

C. Hostnames

D. Labels

**Answer: D**

**Timestamp: March 1, 2025, 3:44 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157332-exam-kcna-topic-1-question-106-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D The nodeSelector field within a PodSpec is used to place Pods on target nodes based on the labels assigned to those nodes. Labels are key-value pairs that can be attached to Kubernetes objects, including nodes, to provide metadata that can be used for selection and organization. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 107 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 107
Topic #: 1

[All KCNA Questions]

What do Deployments and StatefulSets have in common? 
Suggested Answer: A 🗳️ 

A. They manage Pods that are based on an identical container spec.

B. They support the OnDelete update strategy.

C. They support an ordered, graceful deployment and scaling.

D. They maintain a sticky identity for each of their Pods.

**Answer: A**

**Timestamp: March 1, 2025, 3:55 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157333-exam-kcna-topic-1-question-107-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A Both Deployments and StatefulSets in Kubernetes manage Pods that are based on an identical container specification. This means that the Pods managed by a Deployment or a StatefulSet are created from the same template, ensuring consistency in the container images, environment variables, volumes, and other configurations. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 108 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 108
Topic #: 1

[All KCNA Questions]

What is the practice of bringing financial accountability to the variable spend model of cloud resources? 
Suggested Answer: D 🗳️ 

A. FaaS

B. DevOps

C. CloudCost

D. FinOps

**Answer: D**

**Timestamp: March 1, 2025, 3:58 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157336-exam-kcna-topic-1-question-108-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D FinOps (short for Financial Operations) is the practice of bringing financial accountability to the variable spend model of cloud resources. It involves collaboration between finance, operations, and engineering teams to manage and optimize cloud costs while maintaining the agility and scalability benefits of cloud computing. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 109 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 109
Topic #: 1

[All KCNA Questions]

What is a best practice to minimize the container image size? 
Suggested Answer: B 🗳️ 

A. Use a DockerFile.

B. Use multistage builds.

C. Build images with different tags.

D. Add a build.sh script.

**Answer: B**

**Timestamp: March 1, 2025, 3:59 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157337-exam-kcna-topic-1-question-109-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B Use a minimal base image and multi-stage builds. Minimizing the container image size is important for improving build times, reducing attack surfaces, and optimizing resource usag upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 110 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 110
Topic #: 1

[All KCNA Questions]

Which tools enable Kubernetes HorizontalPodAutoscalers to use custom, application-generated metrics to trigger scaling events? 
Suggested Answer: A 🗳️ 

A. Prometheus and the prometheus-adapter.

B. Graylog and graylog-autoscaler metrics.

C. Graylog and the kubernetes-adapter.

D. Grafana and Prometheus.

**Answer: A**

**Timestamp: Jan. 12, 2025, 10:01 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/154446-exam-kcna-topic-1-question-110-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A Prometheus and the prometheus-adapter enable Kubernetes HorizontalPodAutoscalers to use custom, application-generated metrics to trigger scaling events. Prometheus collects the metrics, and the prometheus-adapter exposes them to the Kubernetes API, allowing the HPA to make scaling decisions based on these custom metrics. upvoted 1 times ... 2211094 12 months ago Selected Answer: A A. Prometheus and the prometheus-adapter Prometheus is widely used for monitoring and collecting metrics in Kubernetes environments. The prometheus-adapter allows Kubernetes HorizontalPodAutoscalers (HPAs) to use custom, application-generated metrics from Prometheus to trigger scaling events. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 111 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 111
Topic #: 1

[All KCNA Questions]

Which of the following is a valid PromQL query? 
Suggested Answer: D 🗳️ 

A. SELECT * from http_requests_total WHERE job=apiserver

B. http_requests_total WHERE (job="apiserver")

C. SELECT * from http_requests_total

D. http_requests_total(job="apiserver")

**Answer: D**

**Timestamp: Sept. 23, 2024, 2:39 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/147997-exam-kcna-topic-1-question-111-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D The valid PromQL query to select metrics with a specific label is http_requests_total{job="apiserver"}. This query retrieves all time series with the metric name http_requests_total where the job label is equal to apiserver. upvoted 1 times ... 2211094 12 months ago Selected Answer: D D. http_requests_total{job="apiserver"} In PromQL (Prometheus Query Language), the syntax for filtering metrics based on label values uses curly braces {}. Option D correctly filters the http_requests_total metric to include only the metrics where the label job equals "apiserver". upvoted 1 times ... miskill 1 year, 3 months ago Selected Answer: D D is correct but the right way to write it is: http_requests_total{job="apiserver"} upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 112 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 112
Topic #: 1

[All KCNA Questions]

Which of the following best describes horizontally scaling an application deployment? 
Suggested Answer: C 🗳️ 

A. The act of adding/removing node instances to the cluster to meet demand.

B. The act of adding/removing applications to meet demand.

C. The act of adding/removing application instances of the same application to meet demand.

D. The act of adding/removing resources to application instances to meet demand.

**Answer: C**

**Timestamp: March 1, 2025, 4:15 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157339-exam-kcna-topic-1-question-112-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C Horizontally scaling an application deployment involves adding or removing instances of the same application to meet demand upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 113 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 113
Topic #: 1

[All KCNA Questions]

How many different Kubernetes service types can you define? 
Suggested Answer: C 🗳️ 

A. 2

B. 3

C. 4

D. 5

**Answer: C**

**Timestamp: Sept. 12, 2024, 2:21 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/147404-exam-kcna-topic-1-question-113-discussion/)

Comments: iulian0585 6 months, 2 weeks ago Selected Answer: C A Headless Service is a special case of the ClusterIP type, distinguished by setting spec.clusterIP: None. It does not introduce a new Service type but modifies the behavior of ClusterIP. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: C Kubernetes supports four different Service types: ClusterIP, NodePort, LoadBalancer, and ExternalName. Each Service type serves a specific purpose and use case, allowing you to expose applications running in Pods to other applications or users, both inside and outside the cluster. upvoted 1 times ... Atest00678 11 months, 3 weeks ago Selected Answer: D Kubernetes Services have the following service types: • ClusterIP – default, randomly forward traffic to any pod set with a target port • Headless – send traffic to the specific pod, when you have stateful pods eg. Database • NodePort – external service allows you to use worker node IP address • LoadBalancer – similar to NodePort except leverages Cloud Service Provider’s (CSPs) load balancer • ExternalName — a special service that does not have selectors and uses DNS names instead upvoted 2 times ... miskill 1 year, 3 months ago Selected Answer: C The correct answer is: C. 4 Explanation: Kubernetes supports four different service types: ClusterIP: The default service type, which exposes the service on a cluster-internal IP. It is accessible only within the cluster. NodePort: Exposes the service on each node's IP at a static port. This makes the service accessible outside the cluster via <NodeIP>:<NodePort>. LoadBalancer: Exposes the service externally using a cloud provider's load balancer. This service type automatically provisions a load balancer for the service. ExternalName: Maps the service to a DNS name by returning a CNAME record, redirecting traffic to an external service outside the Kubernetes cluster. These are the four primary service types in Kubernetes. upvoted 4 times ... jafg22 1 year, 3 months ago The correct answer its C ClusterIP: The default type, exposes the Service on a cluster-internal IP. NodePort: Exposes the Service on each Node’s IP at a static port. LoadBalancer: Exposes the Service externally using a cloud provider's load balancer. ExternalName: Maps a Service to an external DNS name. upvoted 2 times ... Darkren4eveR 1 year, 4 months ago c https://kubernetes.io/docs/concepts/services-networking/service/ upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 114 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 114
Topic #: 1

[All KCNA Questions]

What is the difference between a Deployment and a ReplicaSet? 
Suggested Answer: D 🗳️ 

A. With a Deployment, you can’t control the number of pod replicas.

B. A ReplicaSet does not guarantee a stable set of replica pods running.

C. A Deployment is basically the same as a ReplicaSet with annotations.

D. A Deployment is a higher-level concept that manages ReplicaSets.

**Answer: D**

**Timestamp: March 1, 2025, 4:18 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157340-exam-kcna-topic-1-question-114-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D A Deployment is a higher-level abstraction that manages ReplicaSets and provides declarative updates to applications. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 115 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 115
Topic #: 1

[All KCNA Questions]

The Container Runtime Interface (CRI) defines the protocol for the communication between: 
Suggested Answer: A 🗳️ 

A. The kubelet and the container runtime.

B. The container runtime and etcd.

C. The kube-apiserver and the kubelet.

D. The container runtime and the image registry.

**Answer: A**

**Timestamp: March 1, 2025, 4:20 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157342-exam-kcna-topic-1-question-115-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A The Container Runtime Interface (CRI) defines the protocol for communication between the kubelet and the container runtime. The CRI allows Kubernetes to support different container runtimes in a consistent manner by providing a standard interface for the kubelet to interact with the container runtime. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 116 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 116
Topic #: 1

[All KCNA Questions]

Which authorization-mode allows granular control over the operations that different entities can perform on different objects in a Kubernetes cluster? 
Suggested Answer: B 🗳️ 

A. Webhook Mode Authorization Control

B. Role Based Access Control

C. Node Authorization Access Control

D. Attribute Based Access Control

**Answer: B**

**Timestamp: March 1, 2025, 4:22 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157343-exam-kcna-topic-1-question-116-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B Role-Based Access Control (RBAC) is an authorization mode in Kubernetes that provides fine-grained control over the actions that users, groups, and service accounts can perform on various resources within the cluster. RBAC allows administrators to define roles and role bindings to specify permissions for different entities. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 117 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 117
Topic #: 1

[All KCNA Questions]

Which of the following is a correct definition of a Helm chart? 
Suggested Answer: D 🗳️ 

A. A Helm chart is a collection of YAML files bundled in a tar.gz file and can be applied without decompressing it.

B. A Helm chart is a collection of JSON files and contains all the resource definitions to run an application on Kubernetes.

C. A Helm chart is a collection of YAML files that can be applied on Kubernetes by using the kubectl tool.

D. A Helm chart is similar to a package and contains all the resource definitions to run an application on Kubernetes.

**Answer: D**

**Timestamp: March 1, 2025, 4:23 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157344-exam-kcna-topic-1-question-117-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: D A correct definition of a Helm chart is that it is a package or collection of files that defines a related set of Kubernetes resources, using templates and default configurations to manage and deploy a Kubernetes application. Helm charts serve as a template for applications, enabling repeatable installations, upgrades, and rollbacks of services in various environments upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: D A Helm chart is a package format used by Helm, a package manager for Kubernetes. Helm charts contain all the necessary resource definitions to deploy and manage an application on a Kubernetes cluster. These resource definitions are typically written in YAML and include templates for Kubernetes objects such as Deployments, Services, ConfigMaps, and more upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 118 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 118
Topic #: 1

[All KCNA Questions]

Which of the following sentences is true about namespaces in Kubernetes? 
Suggested Answer: C 🗳️ 

A. You can create a namespace within another namespace in Kubernetes.

B. You can create two resources of the same kind and name in a namespace.

C. The default namespace exists when a new cluster is created.

D. All the objects in the cluster are namespaced by default.

**Answer: C**

**Timestamp: March 1, 2025, 4:26 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157345-exam-kcna-topic-1-question-118-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: C The default namespace exists when a new cluster is created. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: C The default namespace exists when a new cluster is created. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 119 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 119
Topic #: 1

[All KCNA Questions]

How does Horizontal Pod autoscaling work in Kubernetes? 
Suggested Answer: D 🗳️ 

A. The Horizontal Pod Autoscaler controller adds more CPU or memory to the pods when the load is above the configured threshold, and reduces CPU or memory when the load is below.

B. The Horizontal Pod Autoscaler controller adds more pods when the load is above the configured threshold, but does not reduce the number of pods when the load is below.

C. The Horizontal Pod Autoscaler controller adds more pods to the specified DaemonSet when the load is above the configured threshold, and reduces the number of pods when the load is below.

D. The Horizontal Pod Autoscaler controller adds more pods when the load is above the configured threshold, and reduces the number of pods when the load is below.

**Answer: D**

**Timestamp: March 1, 2025, 4:29 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157347-exam-kcna-topic-1-question-119-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: D Horizontal Pod Autoscaling (HPA) in Kubernetes automatically adjusts the number of Pod replicas in a Deployment, ReplicaSet, or StatefulSet based on observed resource utilization or custom metrics. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: D The Horizontal Pod Autoscaler (HPA) in Kubernetes automatically adjusts the number of Pods in a Deployment, ReplicaSet, or StatefulSet based on observed metrics such as CPU utilization, memory usage, or custom metrics. The HPA ensures that the number of running Pods matches the current demand, scaling up when the load is high and scaling down when the load is low. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 120 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 120
Topic #: 1

[All KCNA Questions]

What is a Pod? 
Suggested Answer: D 🗳️ 

A. A networked application within Kubernetes.

B. A storage volume within Kubernetes.

C. A single container within Kubernetes.

D. A group of one or more containers within Kubernetes.

**Answer: D**

**Timestamp: March 1, 2025, 4:28 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157346-exam-kcna-topic-1-question-120-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: D A pod has several meanings, most commonly referring to a plant's seed vessel (like a pea pod) or a group of marine animals. It can also mean a small, protective housing for an object, such as an engine pod on an airplane, or in computing, a group of tightly coupled software containers. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: D A group of one or more containers within Kubernetes upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 121 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 121
Topic #: 1

[All KCNA Questions]

What element allows Kubernetes to run Pods across the fleet of nodes? 
Suggested Answer: D 🗳️ 

A. The node server.

B. The etcd static pods.

C. The API server.

D. The kubelet.

**Answer: D**

**Timestamp: March 1, 2025, 4:32 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157349-exam-kcna-topic-1-question-121-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D The kubelet is the primary node agent that runs on each node in a Kubernetes cluster. It is responsible for managing the lifecycle of Pods on its node, including starting, stopping, and monitoring the containers within those Pods. The kubelet ensures that the containers are running as expected and reports the status back to the Kubernetes control plane. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 122 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 122
Topic #: 1

[All KCNA Questions]

What is the Kubernetes object used for running a recurring workload? 
Suggested Answer: D 🗳️ 

A. Job

B. Batch

C. DaemonSet

D. CronJob

**Answer: D**

**Timestamp: March 1, 2025, 4:33 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157350-exam-kcna-topic-1-question-122-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: D A CronJob allows you to schedule a Job to run at specified intervals using cron expressions, similar to the cron utility in Unix-like operating systems. This makes it suitable for tasks such as periodic data backups, scheduled report generation, or other automated, time-based operations. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: D A CronJob is the Kubernetes object used for running a recurring workload upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 123 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 123
Topic #: 1

[All KCNA Questions]

In the DevOps framework and culture, who builds, automates, and offers continuous delivery tools for developer teams? 
Suggested Answer: C 🗳️ 

A. Application Users

B. Application Developers

C. Platform Engineers

D. Cluster Operators

**Answer: C**

**Timestamp: March 1, 2025, 4:35 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157351-exam-kcna-topic-1-question-123-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C In the DevOps framework and culture, Platform Engineers are responsible for building, automating, and offering continuous delivery tools for developer teams. They focus on creating and maintaining the infrastructure and tooling that enable development teams to deliver software efficiently and reliably. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 124 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 124
Topic #: 1

[All KCNA Questions]

Which kubectl command is useful for collecting information about any type of resource that is active in a Kubernetes cluster? 
Suggested Answer: A 🗳️ 

A. describe

B. list

C. expose

D. explain

**Answer: A**

**Timestamp: March 1, 2025, 4:37 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157352-exam-kcna-topic-1-question-124-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: A The primary kubectl commands useful for collecting information about active resources in a Kubernetes cluster are kubectl get and kubectl describe. kubectl get: This command is used to retrieve a summary of information about one or more resources. It provides a quick overview of the resource's status and key attributes. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: A The kubectl describe command is useful for collecting detailed information about any type of resource that is active in a Kubernetes cluster. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 125 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 125
Topic #: 1

[All KCNA Questions]

The cloud native architecture centered around microservices provides a strong system that ensures ______________. 
Suggested Answer: B 🗳️ 

A. fallback

B. resiliency

C. failover

D. high reachability

**Answer: B**

**Timestamp: March 1, 2025, 4:38 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157353-exam-kcna-topic-1-question-125-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B The cloud native architecture centered around microservices provides a strong system that ensures scalability, resilience, and agility. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 126 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 126
Topic #: 1

[All KCNA Questions]

Which of the following is the correct command to run a nginx deployment with 2 replicas? 
Suggested Answer: B 🗳️ 

A. kubectl run deploy nginx --image=nginx --replicas=2

B. kubectl create deploy nginx --image=nginx --replicas=2

C. kubectl create nginx deployment --image=nginx –replicas=2

D. kubectl create deploy nginx --image=nginx --count=2

**Answer: B**

**Timestamp: Oct. 25, 2024, 11:35 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/150237-exam-kcna-topic-1-question-126-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B The kubectl create deploy command is used to create a new Deployment in Kubernetes. The --image flag specifies the container image to use, and the --replicas flag specifies the number of replicas for the Deployment. upvoted 1 times ... heeloco 1 year, 2 months ago Selected Answer: B example: # Create a deployment named my-dep that runs the nginx image with 3 replicas kubectl create deployment my-dep --image=nginx --replicas=3 upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 127 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 127
Topic #: 1

[All KCNA Questions]

What does "Continuous Integration" mean? 
Suggested Answer: B 🗳️ 

A. The continuous integration and testing of code changes from multiple sources manually.

B. The continuous integration and testing of code changes from multiple sources via automation.

C. The continuous integration of changes from one environment to another.

D. The continuous integration of new tools to support developers in a project.

**Answer: B**

**Timestamp: March 1, 2025, 4:40 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157354-exam-kcna-topic-1-question-127-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B Continuous Integration (CI) is a software development practice where developers frequently integrate their code changes into a shared repository, typically multiple times a day. Each integration is automatically verified by running automated tests and builds, ensuring that the codebase remains stable and functional upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 128 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 128
Topic #: 1

[All KCNA Questions]

Which of the following options is true about considerations for large Kubernetes clusters? 
Suggested Answer: C 🗳️ 

A. Kubernetes supports up to 1000 nodes and recommends no more than 1000 containers per node.

B. Kubernetes supports up to 5000 nodes and recommends no more than 500 pods per node.

C. Kubernetes supports up to 5000 nodes and recommends no more than 110 pods per node.

D. Kubernetes supports up to 50 nodes and recommends no more than 1000 containers per node.

**Answer: C**

**Timestamp: Oct. 25, 2024, 11:39 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/150239-exam-kcna-topic-1-question-128-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C Maximum Nodes: Kubernetes supports up to 5000 nodes in a single cluster. Maximum Pods per Node: Kubernetes recommends no more than 110 pods per node. This limit helps ensure that the kubelet and other node components can manage the pods efficiently without performance degradation. upvoted 1 times ... heeloco 1 year, 2 months ago Selected Answer: C - No more than 110 pods per node - No more than 5,000 nodes - No more than 150,000 total pods - No more than 300,000 total containers upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 129 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 129
Topic #: 1

[All KCNA Questions]

Which component of the node is responsible to run workloads? 
Suggested Answer: D 🗳️ 

A. The kubelet.

B. The kubeproxy.

C. The kube-apiserver.

D. The container runtime.

**Answer: D**

**Timestamp: Sept. 15, 2024, 8:54 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/147589-exam-kcna-topic-1-question-129-discussion/)

Comments: mc2301 Highly Voted 1 year, 3 months ago The component of the node responsible for running workloads is the kubelet. While the container runtime (like Docker, containerd, or CRI-O) is crucial for actually running the containers, the kubelet is the orchestrator that ensures the containers are running as specified by the Kubernetes control plane. upvoted 5 times ... iulian0585 Most Recent 6 months, 2 weeks ago Selected Answer: D The correct answer is: D. The container runtime. upvoted 1 times ... donathon 7 months, 3 weeks ago Selected Answer: D The correct answer is: D. The container runtime Explanation: The container runtime is the component of a Kubernetes node that is directly responsible for running workloads, i.e., starting and managing containers. Breakdown of the options: A. The kubelet → ❌ The kubelet manages the node and ensures containers are running, but it delegates the actual running of containers to the container runtime. B. The kube-proxy → ❌ Handles network routing for services, not workloads. C. The kube-apiserver → ❌ Part of the control plane, not a node component; it handles API requests. D. The container runtime → ✅ Runs the containers (e.g., containerd, CRI-O, Docker). Would you like a diagram showing how the kubelet, container runtime, and other node components interact? upvoted 1 times ... Joshua555 9 months, 3 weeks ago Selected Answer: A i think kubelet upvoted 1 times ... shahy0 10 months ago Selected Answer: A B. The kubeproxy Incorrect. kube-proxy is responsible for maintaining network rules on nodes and ensuring that network traffic is correctly routed to the appropriate Pods. It is not responsible for running workloads. C. The kube-apiserver Incorrect. The kube-apiserver is a central component of the Kubernetes control plane that exposes the Kubernetes API. It handles API requests and updates the state of the cluster but does not run workloads. D. The container runtime Incorrect. While the container runtime (e.g., Docker, containerd) is responsible for running containers, it is the kubelet that manages the lifecycle of those containers and ensures they are running according to the Pod specifications. upvoted 1 times ... shahy0 10 months, 2 weeks ago Selected Answer: D The container runtime is the component of the node responsible for running workloads in the form of containers upvoted 1 times ... yoyo2424 11 months, 3 weeks ago Selected Answer: D are D....The container runtime is responsible for running the workloads (containers) on a node in Kubernetes. It pulls container images from a registry and manages the lifecycle of containers. Examples of container runtimes include Docker, containerd, and CRI-O. upvoted 1 times ... 2211094 12 months ago Selected Answer: D Kubelet manage entire life cycle of the containers, while container runtime manages the runtime of the workload. So the answer is D upvoted 1 times ... knob360 1 year, 1 month ago Selected Answer: A The component of the node responsible for running workloads is the kubelet. upvoted 3 times ... 4532da9 1 year, 3 months ago Selected Answer: D The container runtime is responsible for running the containers that make up the workloads on the node. The kubelet interacts with the container runtime to start, stop, and manage containers, but it is the container runtime itself that actually runs the workloads. upvoted 3 times ...
----------------------------------------

## Exam KCNA topic 1 question 130 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 130
Topic #: 1

[All KCNA Questions]

The IPv4/IPv6 dual stack in Kubernetes: 
Suggested Answer: D 🗳️ 

A. Translates an IPv4 request from a service to an IPv6 service.

B. Allows you to access the IPv4 address by using the IPv6 address.

C. Requires NetworkPolicies to prevent services from mixing requests.

D. Allows you to create IPv4 and IPv6 dual stack services.

**Answer: D**

**Timestamp: March 1, 2025, 4:42 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157355-exam-kcna-topic-1-question-130-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D Allows you to create IPv4 and IPv6 dual stack services upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 131 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 131
Topic #: 1

[All KCNA Questions]

What does "continuous" mean in the context of CI/CD? 
Suggested Answer: C 🗳️ 

A. Frequent releases, Manual processes, Repeatable, Fast processing

B. Periodic releases, Manual processes, Repeatable, Automated Processing

C. Frequent releases, Automated processes, Repeatable, Fast processing

D. Periodic releases, Automated processes, Repeatable, Automated processing

**Answer: C**

**Timestamp: March 2, 2025, 5:53 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157419-exam-kcna-topic-1-question-131-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C CI/CD (Continuous Integration and Continuous Deployment/Delivery), "continuous" refers to the practice of frequently integrating code changes, automating the build and deployment processes, ensuring repeatability, and achieving fast processing. This approach aims to improve the efficiency, reliability, and speed of software development and delivery. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 132 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 132
Topic #: 1

[All KCNA Questions]

What is ephemeral storage? 
Suggested Answer: A 🗳️ 

A. Storage space that need not persist across restarts.

B. Storage that may grow dynamically.

C. Storage used by multiple consumers (e.g. multiple Pods).

D. Storage that is always provisioned locally

**Answer: A**

**Timestamp: March 2, 2025, 5:56 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157420-exam-kcna-topic-1-question-132-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A Ephemeral storage refers to temporary storage that is used by containers and does not need to persist across restarts or reboots. This type of storage is typically used for transient data that is only needed for the duration of the container's lifecycle. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 133 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 133
Topic #: 1

[All KCNA Questions]

What is the reference implementation of the OCI runtime specification? 
Suggested Answer: C 🗳️ 

A. lxc

B. cri-o

C. runc

D. docker

**Answer: C**

**Timestamp: March 2, 2025, 5:56 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157421-exam-kcna-topic-1-question-133-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: C runC is the reference implementation of the OCI runtime specification. It is a lightweight, portable container runtime that provides a standardized way to run containers. runC was originally extracted from Docker and has since become the reference implementation for the OCI runtime specification, ensuring compatibility and interoperability across different container runtimes. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 134 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 134
Topic #: 1

[All KCNA Questions]

What is a Service? 
Suggested Answer: B 🗳️ 

A. A static network mapping from a Pod to a port.

B. A way to expose an application running on a set of Pods.

C. The network configuration for a group of Pods.

D. An NGINX load balancer that gets deployed for an application.

**Answer: B**

**Timestamp: March 2, 2025, 5:58 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157422-exam-kcna-topic-1-question-134-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B In Kubernetes, a Service is an abstraction that defines a logical set of Pods and a policy by which to access them. Services provide a stable endpoint (IP address and DNS name) to access a group of Pods, even as the underlying Pods are created, destroyed, or rescheduled. This abstraction allows for reliable communication between different parts of an application or between applications. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 135 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 135
Topic #: 1

[All KCNA Questions]

What's the difference between a security profile and a security context? 
Suggested Answer: B 🗳️ 

A. Security Contexts configure Clusters and Namespaces at runtime. Security profiles are control plane mechanisms to enforce specific settings in the Security Context.

B. Security Contexts configure Pods and Containers at runtime. Security profiles are control plane mechanisms to enforce specific settings in the Security Context.

C. Security Profiles configure Pods and Containers at runtime. Security Contexts are control plane mechanisms to enforce specific settings in the Security Profile.

D. Security Profiles configure Clusters and Namespaces at runtime. Security Contexts are control plane mechanisms to enforce specific settings in the Security Profile.

**Answer: B**

**Timestamp: March 2, 2025, 6:19 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157429-exam-kcna-topic-1-question-135-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B A security context defines the security settings for a Pod or container. It specifies the security-related attributes that should be applied to the Pod or container upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 136 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 136
Topic #: 1

[All KCNA Questions]

At which layer would distributed tracing be implemented in a cloud native deployment? 
Suggested Answer: B 🗳️ 

A. Network

B. Application

C. Database

D. Infrastructure

**Answer: B**

**Timestamp: March 2, 2025, 6:06 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157424-exam-kcna-topic-1-question-136-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: B Distributed tracing is implemented at the application layer in a cloud native deployment. It involves instrumenting application code to capture trace data, propagating trace context across service boundaries, and sending trace data to a tracing backend for analysis. This approach helps in understanding the performance and behavior of distributed applications, identifying bottlenecks, and diagnosing issues. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 137 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 137
Topic #: 1

[All KCNA Questions]

What framework does Kubernetes use to authenticate users with JSON Web Tokens? 
Suggested Answer: A 🗳️ 

A. OpenID Connect

B. OpenID Container

C. OpenID Cluster

D. OpenID CNCF

**Answer: A**

**Timestamp: March 2, 2025, 6:08 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157425-exam-kcna-topic-1-question-137-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: A OpenID Connect (OIDC) is an identity layer built on top of the OAuth 2.0 protocol. It allows clients to verify the identity of users based on the authentication performed by an authorization server, as well as to obtain basic profile information about the user. Kubernetes can use OIDC to authenticate users by validating the JWTs issued by an OIDC-compliant identity provider. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 138 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 138
Topic #: 1

[All KCNA Questions]

Which of the following is a feature Kubernetes provides by default as a container orchestration tool? 
Suggested Answer: D 🗳️ 

A. A portable operating system.

B. File system redundancy.

C. A container image registry.

D. Automated rollouts and rollbacks.

**Answer: D**

**Timestamp: March 2, 2025, 6:10 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157426-exam-kcna-topic-1-question-138-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D Kubernetes can automatically roll out updates to applications and roll back to previous versions if something goes wrong. This ensures that updates are applied smoothly and that the system can recover from failures upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 139 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 139
Topic #: 1

[All KCNA Questions]

Which of the following sentences is true about container runtimes in Kubernetes? 
Suggested Answer: D 🗳️ 

A. If you let iptables see bridged traffic, you don't need a container runtime.

B. If you enable IPv4 forwarding, you don't need a container runtime.

C. Container runtimes are deprecated, you must install CRI on each node

D. You must install a container runtime on each node to run pods on it.

**Answer: D**

**Timestamp: March 2, 2025, 6:18 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157428-exam-kcna-topic-1-question-139-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D Container Runtime: The container runtime is the software that actually runs and manages containers on a node. Examples of container runtimes include Docker, containerd, and CRI-O. The container runtime is responsible for managing the lifecycle of containers, and it is a critical component for running containerized applications in a Kubernetes cluster. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 140 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 140
Topic #: 1

[All KCNA Questions]

If a Pod was waiting for container images to download on the scheduled node, what state would it be in? 
Suggested Answer: D 🗳️ 

A. Failed

B. Succeeded

C. Unknown

D. Pending

**Answer: D**

**Timestamp: March 2, 2025, 6:14 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157427-exam-kcna-topic-1-question-140-discussion/)

Comments: shahy0 10 months, 2 weeks ago Selected Answer: D When a Pod is created in Kubernetes and is waiting for container images to be downloaded on the scheduled node, it will be in the Pending state. The Pending state indicates that the Pod has been accepted by the Kubernetes system, but one or more of the containers in the Pod have not yet been started. This can happen for several reasons, including waiting for the container images to be pulled from the container registry upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 151 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 151
Topic #: 1

[All KCNA Questions]

Which of the following scenarios would benefit the most from a service mesh architecture? 
Suggested Answer: D 🗳️ 

A. A few applications with hundreds of pod replicas running in multiple clusters, each one providing multiple services.

B. Thousands of distributed applications running in a single cluster, each one providing multiple services.

C. Tens of distributed applications running in multiple clusters, each one providing multiple services.

D. Thousands of distributed applications running in multiple clusters, each one providing multiple services.

**Answer: D**

**Timestamp: March 4, 2025, 4:50 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157519-exam-kcna-topic-1-question-151-discussion/)

Comments: shahy0 10 months, 1 week ago Selected Answer: D A service mesh architecture is most beneficial in scenarios with thousands of distributed applications running in multiple clusters, each one providing multiple services. This complexity requires advanced traffic management, security, observability, and resilience features that a service mesh provides, making it an ideal solution for managing and operating microservices at scale. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 152 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 152
Topic #: 1

[All KCNA Questions]

Kubernetes___ protect you against voluntary interruptions (such as deleting Pods, draining nodes) to run applications in a highly available manner. 
Suggested Answer: B 🗳️ 

A. Pod Topology Spread Constraints

B. Pod Disruption Budgets

C. Taints and Tolerances

D. Resource Limits and Requests

**Answer: B**

**Timestamp: Sept. 9, 2025, 8:07 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/312175-exam-kcna-topic-1-question-152-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: B A Pod Disruption Budget (PDB) is a Kubernetes resource that protects applications from disruption by defining the minimum number or percentage of pods that must remain available during voluntary disruptions like node maintenance, rolling updates, or scaling. PDBs ensure a certain level of application availability is maintained by preventing Kubernetes from taking down too many pods simultaneously, which helps to prevent service outages or degraded performance, according to CloudThat upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 155 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 155
Topic #: 1

[All KCNA Questions]

Manual reclamation policy of a PVC resource is known as: 
Suggested Answer: C 🗳️ 

A. claimRef

B. Delete

C. Retain

D. Recycle

**Answer: C**

**Timestamp: May 22, 2025, 9:53 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/304032-exam-kcna-topic-1-question-155-discussion/)

Comments: 8378ffe 3 months, 4 weeks ago Selected Answer: C The policy for manual reclamation of a PersistentVolume (PV) associated with a PersistentVolumeClaim (PVC) is the Retain reclaim policy. When a PVC is deleted with the Retain policy, the PV and its data remain intact but are not available for another claim until an administrator manually removes the PV and cleans up the underlying storage. upvoted 1 times ... fremadse 6 months, 2 weeks ago Selected Answer: C "Retain – Manual reclamation. The PV still exists and must be manually handled by an administrator. The data and PV object are preserved." upvoted 1 times ... iulian0585 6 months, 2 weeks ago Selected Answer: C In Kubernetes, the reclamation policy of a PersistentVolume (PV) determines what happens to the underlying storage resource when the PersistentVolumeClaim (PVC) is deleted. There are three main reclamation policies: Delete – The associated storage (e.g., EBS volume) is deleted automatically. Retain – The volume is not deleted. Manual intervention is required to reclaim or reuse it. Recycle – Deprecated. It used to wipe data and make the volume available again, but it’s no longer recommended. upvoted 1 times ... TheBigMan 7 months, 3 weeks ago Selected Answer: A bcd are policy's but with A you can manualy claim in yaml upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 156 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 156
Topic #: 1

[All KCNA Questions]

Which component of the Kubernetes architecture is responsible for integration with the CRI container runtime? 
Suggested Answer: B 🗳️ 

A. kubeadm

B. kubelet

C. kube-aplserver

D. kubectl

**Answer: B**

**Timestamp: Sept. 9, 2025, 8:25 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/312176-exam-kcna-topic-1-question-156-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: B The component of the Kubernetes architecture responsible for integration with the Container Runtime Interface (CRI) is the kubelet. The kubelet is an agent that runs on each worker node in a Kubernetes cluster. It interacts with the container runtime on that node through the CRI, managing the lifecycle of containers and reporting their status back to the Kubernetes control plane. The CRI provides a standardized interface that allows the kubelet to communicate with various container runtimes (e.g., containerd, CRI-O) without needing to know their specific implementations. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 157 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 157
Topic #: 1

[All KCNA Questions]

Which one of the following is an open source runtime security tool? 
Suggested Answer: C 🗳️ 

A. lxd

B. containerd

C. falco

D. gvisor

**Answer: C**

**Timestamp: March 4, 2025, 4:52 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157520-exam-kcna-topic-1-question-157-discussion/)

Comments: shahy0 10 months, 1 week ago Selected Answer: C A. lxd Incorrect. LXD is a container and virtual machine manager that provides a user experience similar to virtual machines but with the performance and density of containers. It is not specifically a runtime security tool. B. containerd Incorrect. containerd is an industry-standard container runtime that provides core container management capabilities, such as image transfer, container execution, and storage. It is not specifically a runtime security tool. D. gvisor Incorrect. gVisor is a user-space kernel that provides an additional layer of isolation between the application and the host kernel. While it enhances security by isolating containers, it is not specifically a runtime security monitoring tool like Falco. Falco is an open-source runtime security tool designed to monitor and detect anomalous behavior in containerized environments. It provides runtime security by analyzing system calls and other events, using customizable rules to detect suspicious activities, and integrating with various logging and alerting systems for real-time notifications. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 161 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 161
Topic #: 1

[All KCNA Questions]

Which of these is a valid container restart policy? 
Suggested Answer: D 🗳️ 

A. On login

B. On update

C. On start

D. On failure

**Answer: D**

**Timestamp: March 4, 2025, 4:53 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/157521-exam-kcna-topic-1-question-161-discussion/)

Comments: shahy0 10 months, 1 week ago Selected Answer: D In Kubernetes, a container restart policy defines the conditions under which a container should be restarted. There are three valid restart policies for containers in Kubernetes: Always: The container will always be restarted if it stops, regardless of the exit status. OnFailure: The container will be restarted only if it exits with a non-zero exit status (indicating a failure). Never: The container will not be restarted, regardless of the exit status. The valid container restart policies in Kubernetes are Always, OnFailure, and Never. The correct answer, OnFailure, specifies that the container will be restarted only if it exits with a non-zero exit status, indicating a failure. upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 162 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 162
Topic #: 1

[All KCNA Questions]

Which of the following is a challenge derived from running cloud native applications? 
Suggested Answer: B 🗳️ 

A. The operational costs of maintaining the data center of the company.

B. The cost optimization is complex to maintain across different public cloud environments.

C. The lack of different container images available in public image repositories.

D. The lack of services provided by the most common public clouds

**Answer: B**

**Timestamp: July 18, 2025, 3:54 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306770-exam-kcna-topic-1-question-162-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: A A. Final answer. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 163 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 163
Topic #: 1

[All KCNA Questions]

What is the correct hierarchy of Kubernetes components? 
Suggested Answer: C 🗳️ 

A. Containers → Pods → Cluster → Nodes

B. Nodes → Cluster → Containers → Pods

C. Cluster → Nodes → Pods → Containers

D. Pods → Cluster → Containers → Nodes

**Answer: C**

**Timestamp: July 18, 2025, 3:51 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306737-exam-kcna-topic-1-question-163-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: D My answer: D. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 164 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 164
Topic #: 1

[All KCNA Questions]

Which statement about Secrets is correct? 
Suggested Answer: C 🗳️ 

A. A Secret is part of a Pod specification.

B. Secrets data is encrypted with the cluster private key by default.

C. Secrets data is base64 encoded and stored unencrypted by default

D. A Secret can only be used for confidential data

**Answer: C**

**Timestamp: July 18, 2025, 3:51 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306746-exam-kcna-topic-1-question-164-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: B I’m going with B. The other options seem overly complicated or don't fit the requirements. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 165 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 165
Topic #: 1

[All KCNA Questions]

Which mechanism allows extending the Kubernetes API? 
Suggested Answer: B 🗳️ 

A. ConfigMap

B. CustomResourceDefinition

C. MutatingAdmissionWebhookmechamsm

D. Kustomize

**Answer: B**

**Timestamp: July 18, 2025, 3:51 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306738-exam-kcna-topic-1-question-165-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: B Gotta say, B makes the most sense. It directly addresses the problem statement. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 166 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 166
Topic #: 1

[All KCNA Questions]

Which of the following observability data streams would be most useful when desiring to plot resource consumption and predicted future resource exhaustion? 
Suggested Answer: D 🗳️ 

A. stdout

B. Traces

C. Logs

D. Metrics

**Answer: D**

**Timestamp: July 18, 2025, 3:51 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306741-exam-kcna-topic-1-question-166-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: C 100% C. I remember this specific point from the official study guide. upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 167 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 167
Topic #: 1

[All KCNA Questions]

What can be used to create a job that will run at specified times/dates or on a repeating schedule? 
Suggested Answer: D 🗳️ 

A. Job

B. CalenderJob

C. BatchJob

D. CronJob

**Answer: D**

**Timestamp: July 18, 2025, 3:52 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306756-exam-kcna-topic-1-question-167-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: D I'm voting for D. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 168 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 168
Topic #: 1

[All KCNA Questions]

If kubectl is failing to retrieve information from the cluster, where can you find pod logs to troubleshoot? 
Suggested Answer: A 🗳️ 

A. /var/log/pods/

B. ~/.kube/config

C. /var/log/k8s/

D. /etc/kubernetes/

**Answer: A**

**Timestamp: July 18, 2025, 3:52 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306751-exam-kcna-topic-1-question-168-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: A After consulting with Chatgpt, the answer is: A. upvoted 1 times ... Zangi 5 months, 3 weeks ago Selected Answer: B Kinda torn, but B looks most likely to be what the exam is looking for. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 169 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 169
Topic #: 1

[All KCNA Questions]

Which component in Kubernetes is responsible to watch newly created Pods with no assigned node, and selects a node for them to run on? 
Suggested Answer: D 🗳️ 

A. etcd

B. kube controller-manager

C. kube proxy

D. kube scheduler

**Answer: D**

**Timestamp: July 18, 2025, 3:52 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306757-exam-kcna-topic-1-question-169-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: D he Kubernetes component responsible for watching newly created Pods with no assigned node and selecting a node for them to run on is the kube-scheduler. The kube-scheduler is a control plane component that continuously monitors the Kubernetes API server for new Pods that are in a pending state (meaning they have not yet been assigned to a node). When it identifies such a Pod, it then performs a scheduling process to determine the most suitable node for that Pod to run on, considering various factors such as resource requirements, node availability, affinity/anti-affinity rules, and other policy constraints. Finally, it updates the Pod's definition in the API server to bind it to the selected node. upvoted 1 times ... Zangi 5 months, 3 weeks ago Selected Answer: C Gotta say, C makes the most sense. It directly addresses the problem statement. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 170 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 170
Topic #: 1

[All KCNA Questions]

Which control plane component is responsible for updating the node Ready condition if a node becomes unreachable? 
Suggested Answer: B 🗳️ 

A. The kube-oroxy.

B. The node controller.

C. The kubelct.

D. The kube-apiserver

**Answer: B**

**Timestamp: July 18, 2025, 3:50 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306735-exam-kcna-topic-1-question-170-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: B The Node Controller, a component of the Kubernetes control plane, is responsible for updating the Node's Ready condition to Unknown or NotReady when it becomes unreachable and fails to send heartbeats to the control plane within a specific grace period. It continuously monitors node health, and when a node fails to report back, the Node Controller takes action to mark the node as unavailable and initiates the rescheduling of its pods to healthier nodes. upvoted 1 times ... Zangi 5 months, 3 weeks ago Selected Answer: B Feels like B is the intended answer here. The wording is tricky. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 171 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 171
Topic #: 1

[All KCNA Questions]

Which GitOps engine can be used to orchestrate parallel jobs on Kubernetes? 
Suggested Answer: D 🗳️ 

A. Jenkins X

B. Flagger

C. Flux

D. Argo Workflows

**Answer: D**

**Timestamp: July 18, 2025, 3:51 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306743-exam-kcna-topic-1-question-171-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: A 90% sure it's A. The scenario points directly to it. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 172 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 172
Topic #: 1

[All KCNA Questions]

What is the main purpose of the Open Container Initiative (OCI)? 
Suggested Answer: B 🗳️ 

A. Accelerating the adoption of containers and Kubernetes in the industry.

B. Creating open industry standards around container formats and runtimes.

C. Creating industry standards around container formats and runtimes for private purposes.

D. Improving the security of standards around container formats and runtimes

**Answer: B**

**Timestamp: July 18, 2025, 3:53 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306767-exam-kcna-topic-1-question-172-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: B The main purpose of the Open Container Initiative (OCI) is to create and promote universal, vendor-neutral, open industry standards for container formats and runtimes, ensuring interoperability and portability across different platforms and tools. By providing a common baseline, the OCI allows container images to be built by one tool and run by another, fostering a more robust and flexible container ecosystem. upvoted 1 times ... Zangi 5 months, 3 weeks ago Selected Answer: A Absolutely A. The other options are factually incorrect. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 173 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 173
Topic #: 1

[All KCNA Questions]

Which are the core features provided by a service mesh? 
Suggested Answer: A 🗳️ 

A. Authentication and authorization

B. Distributing and replicating data

C. Security vulnerability scanning

D. Configuration management

**Answer: A**

**Timestamp: July 18, 2025, 3:51 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306736-exam-kcna-topic-1-question-173-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: C C. Final answer. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 174 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 174
Topic #: 1

[All KCNA Questions]

Which of the following options include only mandatory fields to create a Kubernetes object using a YAML file? 
Suggested Answer: D 🗳️ 

A. apiVersion, template, kind, status

B. apiVersion, metadata, status, spec

C. apiVersion, template, kind, spec

D. apiVersion, metadata, kind, spec

**Answer: D**

**Timestamp: July 18, 2025, 3:52 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306755-exam-kcna-topic-1-question-174-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: D The mandatory fields required to create a Kubernetes object using a YAML file are: apiVersion: Specifies the version of the Kubernetes API you are using to create the object. kind: Defines the type of Kubernetes object you want to create (e.g., Pod, Deployment, Service). metadata: Contains data that helps uniquely identify the object, such as its name and optional namespace. spec: Describes the desired state of the object, varying significantly depending on the kind of object. upvoted 1 times ... Zangi 5 months, 3 weeks ago Selected Answer: B Feels like B is the one. Does anyone disagree? upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 175 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 175
Topic #: 1

[All KCNA Questions]

Which of the following is the name of a container orchestration software? 
Suggested Answer: C 🗳️ 

A. OpenStack

B. Docker

C. Apache Mesos

D. CRI-O

**Answer: C**

**Timestamp: July 18, 2025, 3:52 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306747-exam-kcna-topic-1-question-175-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: C Is C the consensus here? It seems right to me. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 176 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 176
Topic #: 1

[All KCNA Questions]

What happens with a regular pod running in Kubernetes when a node fails? 
Suggested Answer: B 🗳️ 

A. A new pod with the same UID is scheduled to another node after a while.

B. A new, near-identical pod but with different UID is scheduled to another node.

C. By default, a pod can only be scheduled to the same node when the node fails.

D. A new pod is scheduled on a different node only if it is configured explicitly.

**Answer: B**

**Timestamp: July 18, 2025, 3:53 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306763-exam-kcna-topic-1-question-176-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: D 90% sure it's D. The scenario points directly to it. upvoted 2 times ...
----------------------------------------

## Exam KCNA topic 1 question 177 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 177
Topic #: 1

[All KCNA Questions]

What is the minimum of etcd members that are required for a highly available Kubernetes cluster? 
Suggested Answer: D 🗳️ 

A. Two etcd members.

B. Five etcd members.

C. Six etcd members.

D. Three etcd members

**Answer: D**

**Timestamp: July 18, 2025, 3:53 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306762-exam-kcna-topic-1-question-177-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: D For a highly available Kubernetes cluster, a minimum of three etcd members is required. Explanation: etcd uses a quorum mechanism to ensure data consistency and prevent split-brain scenarios. A quorum is defined as a strict majority of members required to agree on a decision before it is committed. For a cluster with n members, the quorum is (n/2) + 1. upvoted 2 times ... Zangi 5 months, 3 weeks ago Selected Answer: C C is the only option that isn't immediately wrong. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 178 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 178
Topic #: 1

[All KCNA Questions]

What is the main purpose of the Ingress in Kubernetes? 
Suggested Answer: D 🗳️ 

A. Access HTTP and HTTPS services running in the cluster based on their IP address.

B. Access services different from HTTP or HTTPS running in the cluster based on their IP address.

C. Access services different from HTTP or HTTPS running in the cluster based on their path

D. Access HTTP and HTTPS services running in the cluster based on their path.

**Answer: D**

**Timestamp: July 18, 2025, 3:51 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306739-exam-kcna-topic-1-question-178-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: D Ingress is used to manage incoming traffic to the services in a Kubernetes cluster, acting as an API object that manages external access to services within the cluster. upvoted 1 times ... Zangi 5 months, 3 weeks ago Selected Answer: D My gut says D. It aligns with best practices. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 179 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 179
Topic #: 1

[All KCNA Questions]

How can you extend the Kubernetes API? 
Suggested Answer: A 🗳️ 

A. Adding a CustomResourceDefinition or implementing an aggregation layer.

B. Adding a new version of a resource, for instance v4beta3.

C. With the command kubectl extend api, logged in as an administrator.

D. Adding the desired API object as a kubelet parameter

**Answer: A**

**Timestamp: July 18, 2025, 3:53 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306765-exam-kcna-topic-1-question-179-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: A In summary: CRDs ======= provide a lightweight way to introduce new resource types directly into the existing Kubernetes API server. =================== API Aggregation allows for more complex and independent API extensions by integrating external API servers. upvoted 2 times ... Zangi 5 months, 3 weeks ago Selected Answer: D Leaning toward D, but I'd be interested to hear arguments for the others. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 180 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 180
Topic #: 1

[All KCNA Questions]

What is an ephemeral container? 
Suggested Answer: B 🗳️ 

A. A specialized container that runs as root for infosec applications.

B. A specialized container that runs temporarily in an existing Pod.

C. A specialized container that extends and enhances the ‘main’ container in a Pod

D. A specialized container that run before app container in a Pod

**Answer: B**

**Timestamp: July 18, 2025, 3:52 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306754-exam-kcna-topic-1-question-180-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: B A special type of container that runs temporarily in an existing Pod to accomplish user-initiated actions such as troubleshooting. upvoted 1 times ... Zangi 5 months, 3 weeks ago Selected Answer: B Not super confident, but I’ll go with B. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 181 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 181
Topic #: 1

[All KCNA Questions]

In a cloud native environment, who is usually responsible for maintaining the workloads running across the different platforms? 
Suggested Answer: B 🗳️ 

A. The cloud provider.

B. The Site Reliability Engineering (SRE) team.

C. The team of developers.

D. The Support Engineering team (SE).

**Answer: B**

**Timestamp: July 18, 2025, 3:53 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306761-exam-kcna-topic-1-question-181-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: A I'm arriving at A by process of elimination. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 182 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 182
Topic #: 1

[All KCNA Questions]

Which Kubernetes-native deployment strategy supports zero-downtime updates of a workload? 
Suggested Answer: D 🗳️ 

A. Canary

B. Recreate

C. BlueGreen

D. RollingUpdate

**Answer: D**

**Timestamp: July 18, 2025, 3:52 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306748-exam-kcna-topic-1-question-182-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: D D!!! for sure! The RollingUpdate deployment strategy is a Kubernetes-native method that supports zero-downtime updates by gradually replacing old pods with new ones, ensuring continuous availability during the deployment process. Other strategies that also achieve zero-downtime include Blue-Green Deployment, which uses two identical environments to switch traffic over, and Canary Releases, where updates are first rolled out to a small subset of users. RollingUpdate (Default Strategy) upvoted 1 times ... Zangi 5 months, 3 weeks ago Selected Answer: D My best guess is D right now, but I need to study this area more. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 183 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 183
Topic #: 1

[All KCNA Questions]

What service account does a Pod use in a given namespace when the service account is not specified? 
Suggested Answer: D 🗳️ 

A. admin

B. sysadmin

C. root

D. default

**Answer: D**

**Timestamp: July 18, 2025, 3:51 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306745-exam-kcna-topic-1-question-183-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: D When a Pod is created in a given namespace and no specific service account is explicitly defined in its manifest, Kubernetes automatically assigns the default service account for that particular namespace to the Pod. Every Kubernetes namespace is initialized with a default service account. This default service account is used as a fallback for any Pods that do not specify a serviceAccountName in their spec. The default service account, by default, has no special permissions beyond what Kubernetes grants to all authenticated principals for API discovery, assuming Role-Based Access Control (RBAC) is enabled in the cluster. If a Pod requires specific permissions to interact with the Kubernetes API or other resources, a dedicated service account with the necessary roles and role bindings should be created and explicitly assigned to the Pod. upvoted 1 times ... Zangi 5 months, 3 weeks ago Selected Answer: B My gut says B. It aligns with best practices. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 184 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 184
Topic #: 1

[All KCNA Questions]

What is a cloud native application? 
Suggested Answer: B 🗳️ 

A. It is a monolithic application that has been containerized and is running now on the cloud.

B. It is an application designed to be scalable and take advantage of services running on the cloud.

C. It is an application designed to run all its functions in separate containers.

D. It is any application that runs in a cloud provider and uses its services.

**Answer: B**

**Timestamp: July 18, 2025, 3:53 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306758-exam-kcna-topic-1-question-184-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: B A cloud-native application is software built to run on cloud computing infrastructure, leveraging its scalability, flexibility, and dynamic nature to deliver faster updates and greater agility. These applications are often designed as microservices, which are small, independent components that are easier to update and manage than traditional monolithic applications. Key technologies like containers and orchestration tools are used for automation, while DevOps practices and CI/CD pipelines automate the building, testing, and deployment process upvoted 2 times ... Zangi 5 months, 3 weeks ago Selected Answer: C I’m going with C because the others are clearly non-starters. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 185 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 185
Topic #: 1

[All KCNA Questions]

What's the most adopted way of conflict resolution and decision-making for the open-source projects under the CNCF umbrella? 
Suggested Answer: B 🗳️ 

A. Financial Analysis

B. Discussion and Voting

C. Flipism Technique

D. Project Founder Say

**Answer: B**

**Timestamp: July 18, 2025, 3:53 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306764-exam-kcna-topic-1-question-185-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: B Locked in on B. The other choices are clearly distractors. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 186 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 186
Topic #: 1

[All KCNA Questions]

Which of the following options include resources cleaned by the Kubernetes garbage collection mechanism? 
Suggested Answer: D 🗳️ 

A. Stale or expired CertificateSigningRequests (CSRs) and old deployments.

B. Nodes deleted by a cloud controller manager and obsolete logs from the kubelet.

C. Unused container and container images, and obsolete logs from the kubelet.

D. Terminated pods, completed jobs, and objects without owner references.

**Answer: D**

**Timestamp: July 18, 2025, 3:53 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306760-exam-kcna-topic-1-question-186-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: B I've seen something similar before — pretty sure it was B. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 187 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 187
Topic #: 1

[All KCNA Questions]

What is the default eviction timeout when the Ready condition of a node is Unknown or False? 
Suggested Answer: D 🗳️ 

A. Thirty seconds.

B. Thirty minutes.

C. One minute.

D. Five minutes.

**Answer: D**

**Timestamp: July 18, 2025, 3:52 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306752-exam-kcna-topic-1-question-187-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: C Pretty sure it's C. This pattern is very common. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 188 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 188
Topic #: 1

[All KCNA Questions]

What is the main purpose of a DaemonSet? 
Suggested Answer: A 🗳️ 

A. A DaemonSet ensures that all (or certain) nodes run a copy of a Pod.

B. A DaemonSet ensures that the kubelet is constantly up and running.

C. A DaemonSet ensures that there are as many pods running as specified in the replicas field.

D. A DaemonSet ensures that a process (agent) runs on every node.

**Answer: A**

**Timestamp: July 18, 2025, 3:51 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306740-exam-kcna-topic-1-question-188-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: A A!! The main purpose of a Kubernetes DaemonSet is to ensure that a copy of a specified pod runs on all (or selected) nodes in a cluster. DaemonSets are ideal for deploying background services or node-local tools, such as logging agents, monitoring tools, or network helpers, that need to be present on every node to provide essential cluster functionality. The DaemonSet controller automatically adds the designated pod to new nodes joining the cluster and removes it from nodes that are removed, maintaining the desired state across the cluster upvoted 1 times ... Zangi 5 months, 3 weeks ago Selected Answer: D Absolutely D. The other options are factually incorrect. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 189 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 189
Topic #: 1

[All KCNA Questions]

Why do administrators need a container orchestration tool? 
Suggested Answer: A 🗳️ 

A. To manage the lifecycle of an elevated number of containers.

B. To assess the security risks of the container images used in production.

C. To learn how to transform monolithic applications into microservices.

D. Container orchestration tools such as Kubernetes are the future.

**Answer: A**

**Timestamp: July 18, 2025, 3:52 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306753-exam-kcna-topic-1-question-189-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: B I’m going with B because the others are clearly non-starters. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 190 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 190
Topic #: 1

[All KCNA Questions]

Which two elements are shared between containers in the same pod? 
Suggested Answer: C 🗳️ 

A. Network resources and liveness probes.

B. Storage and container image registry.

C. Storage and network resources.

D. Network resources and Dockerfiles.

**Answer: C**

**Timestamp: July 18, 2025, 3:52 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306749-exam-kcna-topic-1-question-190-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: A Feels like A is the intended answer here. The wording is tricky. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 191 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 191
Topic #: 1

[All KCNA Questions]

In Kubernetes, if the API version of feature is v2beta3, it means that: 
Suggested Answer: B 🗳️ 

A. The version will remain available for all future releases within a Kubernetes major version.

B. The API may change in incompatible ways in a later software release without notice.

C. The software is well tested. Enabling a feature is considered safe.

D. The software may contain bugs. Enabling a feature may expose bugs.

**Answer: B**

**Timestamp: July 18, 2025, 3:52 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306750-exam-kcna-topic-1-question-191-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: C This feels like a 50/50 between two options. I'll go with C. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 192 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 192
Topic #: 1

[All KCNA Questions]

What is the API that exposes resource metrics from the metrics-server? 
Suggested Answer: C 🗳️ 

A. custom.k8s.io

B. resources.k8s.io

C. metrics.k8s.io

D. cadvisor.k8s.io

**Answer: C**

**Timestamp: July 18, 2025, 3:54 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306768-exam-kcna-topic-1-question-192-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: C C!! The API that exposes resource metrics from the Kubernetes metrics-server is the Metrics API, specifically the metrics.k8s.io/v1beta1 API group. This API exposes two primary resource types: NodeMetrics: Provides resource usage information for individual nodes. PodMetrics: Provides resource usage information for individual pods. These metrics are accessible through the Kubernetes API server and are consumed by components like the Horizontal Pod Autoscaler (HPA) and Vertical Pod Autoscaler (VPA), as well as tools like kubectl top. The metrics-server collects this data from the Kubelets running on each node and exposes it through this standardized API. upvoted 1 times ... Zangi 5 months, 3 weeks ago Selected Answer: B The answer is B. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 193 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 193
Topic #: 1

[All KCNA Questions]

Which of the following resources helps in managing a stateless application workload on a Kubernetes cluster? 
Suggested Answer: D 🗳️ 

A. DaemonSet

B. StatefulSet

C. kubectl

D. Deployment

**Answer: D**

**Timestamp: July 18, 2025, 3:54 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306769-exam-kcna-topic-1-question-193-discussion/)

Comments: 8378ffe 4 months ago Selected Answer: D D!! The primary Kubernetes resource for managing a stateless application workload is a Deployment. Deployments are designed for applications where Pods are interchangeable and do not require persistent storage or a stable identity. They manage the desired state of your application, handling tasks like rolling updates, rollbacks, and scaling the number of replica Pods. upvoted 1 times ... Zangi 5 months, 3 weeks ago Selected Answer: D It can't be [Option X] because of [reason]. It can't be [Option Y] because... So it must be D. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 194 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 194
Topic #: 1

[All KCNA Questions]

Which mechanism can be used to automatically adjust the amount of resources for an application? 
Suggested Answer: D 🗳️ 

A. Horizontal Pod Autoscaler (HPA)

B. Kubernetes Event-driven Autoscaling (KEDA)

C. Cluster Autoscaler

D. Vertical Pod Autoscaler (VPA)

**Answer: D**

**Timestamp: June 9, 2025, 3:50 p.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/304639-exam-kcna-topic-1-question-194-discussion/)

Comments: yast89 6 months, 2 weeks ago Selected Answer: D Horizontal Pod Autoscaler : Adjusts the number of Pods. Vertical Pod Autoscaler: Adjusts the resource requests and limits of containers. Kubernetes Event-driven Autoscaling: Provides event-driven autoscaling. Cluster Autoscaler: Adjusts the number of nodes in the cluster. upvoted 1 times ... snag5594 7 months ago Selected Answer: D Its VPA. HPA - auto scales the number of instances (pods) VPA - Auto adjusts the resources upvoted 1 times ... TheBigMan 7 months ago Selected Answer: D ik zeg D upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 195 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 195
Topic #: 1

[All KCNA Questions]

Which of the following is a recommended security habit in Kubernetes? 
Suggested Answer: B 🗳️ 

A. Run the containers as the user with group ID 0 (root) and any user ID.

B. Disallow privilege escalation from within a container as the default option.

C. Run the containers as the user with user ID 0 (root) and any group ID.

D. Allow privilege escalation from within a container as the default option.

**Answer: B**

**Timestamp: July 18, 2025, 3:51 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306742-exam-kcna-topic-1-question-195-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: D I'm gonna take a wild guess and say D. This topic is my weak spot. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 196 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 196
Topic #: 1

[All KCNA Questions]

What are the 3 pillars of Observability? 
Suggested Answer: A 🗳️ 

A. Metrics, Logs, and Traces

B. Metrics, Logs, and Spans

C. Metrics, Data, and Traces

D. Resources, Logs, and Tracing

**Answer: A**

**Timestamp: July 18, 2025, 3:50 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306734-exam-kcna-topic-1-question-196-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: B I’m going with B because the others are clearly non-starters. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 197 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 197
Topic #: 1

[All KCNA Questions]

What edge and service proxy tool is designed to be integrated with cloud native applications? 
Suggested Answer: D 🗳️ 

A. CoreDNS

B. CNI

C. gRPC

D. Envoy

**Answer: D**

**Timestamp: July 18, 2025, 3:53 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306759-exam-kcna-topic-1-question-197-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: D 90% sure it's D. The scenario points directly to it. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 198 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 198
Topic #: 1

[All KCNA Questions]

What is Flux constructed with? 
Suggested Answer: B 🗳️ 

A. GitLab Environment Toolkit

B. GitOps Toolkit

C. Helm Toolkit

D. GitHub Actions Toolkit

**Answer: B**

**Timestamp: July 18, 2025, 3:51 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306744-exam-kcna-topic-1-question-198-discussion/)

Comments: Zangi 5 months, 3 weeks ago Selected Answer: C I've seen something similar before — pretty sure it was C. upvoted 1 times ...
----------------------------------------

## Exam KCNA topic 1 question 199 discussion

Actual exam question from

Linux Foundation's
KCNA

Question #: 199
Topic #: 1

[All KCNA Questions]

In Kubernetes, which abstraction defines a logical set of Pods and a policy by which to access them? 
Suggested Answer: C 🗳️ 

A. Service Account

B. NetworkPolicy

C. Service

D. Custom Resource Definition

**Answer: C**

**Timestamp: July 18, 2025, 3:53 a.m.**

[View on ExamTopics](https://www.examtopics.com/discussions/linux-foundation/view/306766-exam-kcna-topic-1-question-199-discussion/)

Comments: 8378ffe 3 months, 4 weeks ago Selected Answer: A In Kubernetes, the abstraction that defines a logical set of Pods and a policy by which to access them is called a Service. A Service provides a stable network endpoint (IP address and DNS name) for a group of Pods, even as those Pods are created, deleted, or scaled. It abstracts away the ephemeral nature of Pods, allowing other parts of your application to consistently communicate with a set of Pods that perform a specific function, such as a web server or a database. The set of Pods targeted by a Service is typically determined by a selector, which matches labels applied to the Pods. upvoted 1 times ... Zangi 5 months, 3 weeks ago Selected Answer: C I’m going with C. The other options seem overly complicated or don't fit the requirements. upvoted 1 times ...
----------------------------------------

