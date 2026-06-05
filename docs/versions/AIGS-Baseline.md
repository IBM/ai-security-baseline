# Open Source AI Project Governance and Security Baseline

Version: 2026-03-06

{: .warning}
Not for production use.


<button onclick="toTop()" id="topButton" title="Go to top"
style="display: none; position: fixed; bottom: 20px; right: 30px; border: none; background-color: CornflowerBlue; color: white; cursor: pointer; padding: 10px; border-radius: 10px; font-size: 18px;">to top</button>

<script>
let topButton = document.getElementById("topButton");
window.onscroll = function() {scrollFunction()};

function scrollFunction() {
  if (document.documentElement.scrollTop > 50 ) {
    topButton.style.display = "block";
  } else {
    topButton.style.display = "none";
  }
}

function toTop() {
  document.documentElement.scrollTop = 0;
}
</script>


* Contents
{:toc}

## Overview

The Open Source AI Project Governance and Security Baseline (AIGS Baseline) is designed to act as a minimum set of requirements for AI projects relative to its maturity level. It extends the principles of the [OpenSSF Security Baseline](https://baseline.openssf.org/) to address the unique challenges of developing, deploying, and managing Artificial Intelligence (AI) systems. It is designed to be a foundational guide for ensuring AI systems are secure, robust, transparent, and aligned with governance objectives.

For more information on the motive and purpose, see the [FAQ](FAQ.md).

For more information on the project and to make contributions, visit the [GitHub repo](https://github.com/ossf/security-baseline).

---

## Controls Overview

* [Level 1](#level-1): Foundational requirements for all open source AI projects.
* [Level 2](#level-2): Intermediate requirements for maturing AI projects.
* [Level 3](#level-3): Advanced requirements for production AI systems.


### Level 1

**[AIGS-GA-01-01](#aigs-ga-01-01)**: The project MUST document the roles and responsibilities for maintainers
and contributors to the AI project lifecycle from development to
deployment and ongoing monitoring.


**[AIGS-GA-02-01](#aigs-ga-02-01)**: The project MUST provide written, documented disclosure of the AI models ingested in, called, or otherwise used in the development or deployment of the open source AI project.

**[AIGS-GA-02-02](#aigs-ga-02-02)**: Minimum disclosure requirements MUST include the model name, developer, release date, license, model description, deployment status (as necessary), and URL to the open source model repository (if available) or model release announcement.

**[AIGS-GA-03-01](#aigs-ga-03-01)**: The project MUST disclose datasets ingested in or otherwise used in the development or deployment of the open source AI project. This includes datasets used for training, testing, and validation.

**[AIGS-GA-03-02](#aigs-ga-03-02)**: Minimum disclosure requirements for each dataset MUST include the name, country of origin, license, description, data processing (if any), and public URL (if available).

**[AIGS-GA-04-01](#aigs-ga-04-01)**: While active, the project MUST specify a license for the model weights and model files.

**[AIGS-GA-05-01](#aigs-ga-05-01)**: The project MUST define and document the intended use of the open source AI project, including the ethical principles that guide open source AI project development and deployment.

**[AIGS-GA-06-01](#aigs-ga-06-01)**: The project MUST maintain an inventory of its components, tooling, systems and applications, disclosed within the README of the repository where the project is made available.

**[AIGS-SC-01-01](#aigs-sc-01-01)**: The project MUST implement controls to detect and prevent the unintentional sourcing of poisoned data. Controls MUST be documented and disclosed.

**[AIGS-SC-02-01](#aigs-sc-02-01)**: To the extent web crawling is used to source data, web crawling activity MUST respect robots.txt requirements, the crawler MUST be clearly identified, and web crawling activity MUST be disclosed in project documentation.

**[AIGS-SC-03-01](#aigs-sc-03-01)**: The project MUST use a version control system and have a documented process for reviewing and approving changes.

**[AIGS-DI-01-01](#aigs-di-01-01)**: Track the origin and lineage of all data used for training and testing AI models.

**[AIGS-DI-02-01](#aigs-di-02-01)**: The project MUST implement robust security controls to protect the confidentiality, integrity, and availability of data.

**[AIGS-DI-02-02](#aigs-di-02-02)**: Minimum data security requirements MUST include authentication and authorization.

**[AIGS-DI-03-01](#aigs-di-03-01)**: Data access controls MUST be clearly documented and follow the principles of least privilege access.

**[AIGS-DI-04-01](#aigs-di-04-01)**: The project MUST implement controls and testing to prevent the project from inadvertently revealing sensitive information (e.g., PII, proprietary data) from its training set.

**[AIGS-MR-01-01](#aigs-mr-01-01)**: The project MUST undertake red teaming to screen for vulnerabilities to adversarial attacks, including prompt injection attacks.

**[AIGS-MR-01-02](#aigs-mr-01-02)**: To the extent possible, the project owner SHOULD disclose red teaming methodology and sufficient detail to satisfy downstream users of the model&#39;s robustness to a diverse set of adversarial attacks under deployment-like conditions.

**[AIGS-MR-02-01](#aigs-mr-02-01)**: Prior to deployment, the project team MUST test against publicly available or manually-created, task-specific benchmarks in order to evaluate, quantify, and validate the resilience of the AI project against common vulnerabilities.

**[AIGS-MR-02-02](#aigs-mr-02-02)**: To the extent possible, the project owner SHOULD disclose evaluation results and sufficient detail to satisfy downstream users of the model&#39;s reliability under a diverse set of deployment-like conditions and tasks.

**[AIGS-DE-01-01](#aigs-de-01-01)**: The project MUST implement guardrails to mitigate identified risks, and at a minimum, implement guardrails to mitigate against prompt injection attacks (input).

**[AIGS-DE-02-01](#aigs-de-02-01)**: The project MUST have a defined communication channel or mechanism through which vulnerabilities and other security-related information MAY be disclosed to the project contributors and maintainers.

**[AIGS-DE-03-01](#aigs-de-03-01)**: The project MUST validate and sanitize model outputs to prevent downstream vulnerabilities.

**[AIGS-TE-01-01](#aigs-te-01-01)**: Where feasible, the project MUST provide explanations for design decisions that affect project performance and security.

**[AIGS-TE-02-01](#aigs-te-02-01)**: Project contributors MUST maintain clear and comprehensive documentation for the AI system, including its intended use, limitations, and potential risks.


### Level 2

**[AIGS-GA-07-01](#aigs-ga-07-01)**: The project MUST define an AI policy with principles that guide all activities of the organization related to AI and processes for handling deviations and exceptions to policy.

**[AIGS-GA-07-02](#aigs-ga-07-02)**: The AI policy SHOULD consider topic-specific aspects where necessary to provide additional guidance or provide cross-references to other security and governance policies dealing with these core aspects: AI resources and assets, AI system impact assessment scope, and AI project secure development.

**[AIGS-SC-04-01](#aigs-sc-04-01)**: Embedded AI components MUST be scanned as part of development workflows, including CI/CD pipelines, to prevent application security risks from malicious or otherwise compromised components.

**[AIGS-SC-05-01](#aigs-sc-05-01)**: The project MUST document and secure the training pipeline used to produce model artifacts, including dependencies, training code, and environment configuration.

**[AIGS-SC-06-01](#aigs-sc-06-01)**: The project MUST maintain provenance records linking model artifacts, training datasets, training code, and training parameters.

**[AIGS-SC-07-01](#aigs-sc-07-01)**: Where feasible, the project SHOULD support reproducible model builds such that independent parties can recreate the model artifacts from disclosed inputs.

**[AIGS-DI-05-01](#aigs-di-05-01)**: Datasets used for training and evaluation MUST be tracked and versioned to ensure traceability and reproducibility.

**[AIGS-DI-06-01](#aigs-di-06-01)**: The project MUST assess datasets for, at a minimum, legal and copyright risk, licensing compatibility with project intended use, and security risks (e.g., data poisoning, PII/SPI).

**[AIGS-DI-07-01](#aigs-di-07-01)**: The project MUST implement controls to limit the use of unnecessary or high-risk data.

**[AIGS-MR-03-01](#aigs-mr-03-01)**: The project MUST continuously implement red teaming and evaluations prior to pushing system updates or model re-training into production.

**[AIGS-MR-04-01](#aigs-mr-04-01)**: The project MUST maintain regression tests for safety and security behaviors to prevent regression during updates.

**[AIGS-MR-05-01](#aigs-mr-05-01)**: The project MUST document the scope and limitations of testing.

**[AIGS-DE-04-01](#aigs-de-04-01)**: Access to models, tools and tool calling functionality, and general deployment scaffolding or infrastructure MUST follow least-privilege principles.

**[AIGS-DE-05-01](#aigs-de-05-01)**: The project MUST implement logging and monitoring for deployment environments, including tool use, to detect misuse or anomalies.

**[AIGS-DE-06-01](#aigs-de-06-01)**: The project MUST define procedures for responding to security incidents involving the AI project in a timely manner.

**[AIGS-DE-07-01](#aigs-de-07-01)**: The project MUST document processes for releasing model updates and security patches.


### Level 3

**[AIGS-GA-08-01](#aigs-ga-08-01)**: To the extent the open source AI project meets regulatory requirements and standards for the deployment of AI models and systems, the project MUST disclose (non-)compliance.

**[AIGS-GA-09-01](#aigs-ga-09-01)**: While active, structured disclosures of relevant model and data metadata relevant to AI project development and deployment MUST be made available in machine-readable format to inform downstream risk-based controls.

**[AIGS-SC-08-01](#aigs-sc-08-01)**: The project MUST sign AI models, model files, and/or other release artifacts with cryptographic tools and verify them before deployment to prevent the introduction of untrusted components.

**[AIGS-DI-08-01](#aigs-di-08-01)**: Read or write access to data MUST be logged, with timestamp, user id, action performed (read/write), and dataset id and version(s) accessed.

**[AIGS-MR-06-01](#aigs-mr-06-01)**: The project MUST conduct or commission independent testing for adversarial robustness from trusted third-parties.

**[AIGS-MR-07-01](#aigs-mr-07-01)**: The project MUST evaluate models used in the AI project at runtime against an evolving suite of standardized robustness benchmarks.

**[AIGS-DE-08-01](#aigs-de-08-01)**: Deployment environments MUST include runtime safeguards for detecting and mitigating malicious inputs and outputs and unexpected or unauthorized system-level behavior.

**[AIGS-DE-09-01](#aigs-de-09-01)**: The project MUST implement a security mechanism designed to immediately halt, disable, or contain an AI project if it is believed to behave in dangerous or unpredictable ways causing harm.

**[AIGS-TE-03-01](#aigs-te-03-01)**: The project MUST document residual risks and known limitations of safeguards.

**[AIGS-TE-04-01](#aigs-te-04-01)**: The project MUST maintain a public record of safety and security incidents and ex post mitigations.




## Governance and Accountability

Controls related to organizational governance, roles, policies, and accountability for AI projects.



### AIGS-GA-01 - Contributor Roles and Responsibilities

Establish clear documentation of roles and responsibilities for maintainers and contributors throughout the AI project lifecycle from development to deployment and ongoing monitoring.



#### AIGS-GA-01-01

**Requirement:** The project MUST document the roles and responsibilities for maintainers and contributors to the AI project lifecycle from development to deployment and ongoing monitoring.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-GA-02 - Model Governance

Ensure transparency through written, documented disclosure of AI models ingested in, called, or otherwise used in the development or deployment of the open source AI project.



#### AIGS-GA-02-01

**Requirement:** The project MUST provide written, documented disclosure of the AI models ingested in, called, or otherwise used in the development or deployment of the open source AI project.



**Control applies to:**
- Level 1
- Level 2
- Level 3


#### AIGS-GA-02-02

**Requirement:** Minimum disclosure requirements MUST include the model name, developer, release date, license, model description, deployment status (as necessary), and URL to the open source model repository (if available) or model release announcement.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-GA-03 - Data Governance

Achieve transparency by disclosing datasets ingested in or otherwise used in the development or deployment of the open source AI project, including datasets used for training, testing, and validation.



#### AIGS-GA-03-01

**Requirement:** The project MUST disclose datasets ingested in or otherwise used in the development or deployment of the open source AI project. This includes datasets used for training, testing, and validation.



**Control applies to:**
- Level 1
- Level 2
- Level 3


#### AIGS-GA-03-02

**Requirement:** Minimum disclosure requirements for each dataset MUST include the name, country of origin, license, description, data processing (if any), and public URL (if available).



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-GA-04 - Licensing

Establish clear licensing by specifying a license for the model weights and model files while the project is active.



#### AIGS-GA-04-01

**Requirement:** While active, the project MUST specify a license for the model weights and model files.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-GA-05 - Ethics and Intended Use

Define and document the intended use of the open source AI project, including the ethical principles that guide open source AI project development and deployment.



#### AIGS-GA-05-01

**Requirement:** The project MUST define and document the intended use of the open source AI project, including the ethical principles that guide open source AI project development and deployment.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-GA-06 - Project Inventory

Maintain comprehensive visibility through an inventory of components, tooling, systems and applications, disclosed within the README of the repository where the project is made available.



#### AIGS-GA-06-01

**Requirement:** The project MUST maintain an inventory of its components, tooling, systems and applications, disclosed within the README of the repository where the project is made available.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-GA-07 - AI Policy

Establish governance framework by defining an AI policy with principles that guide all activities of the organization related to AI and processes for handling deviations and exceptions to policy.



#### AIGS-GA-07-01

**Requirement:** The project MUST define an AI policy with principles that guide all activities of the organization related to AI and processes for handling deviations and exceptions to policy.



**Control applies to:**
- Level 2
- Level 3


#### AIGS-GA-07-02

**Requirement:** The AI policy SHOULD consider topic-specific aspects where necessary to provide additional guidance or provide cross-references to other security and governance policies dealing with these core aspects: AI resources and assets, AI system impact assessment scope, and AI project secure development.


**Recommendation:** Address AI resources and assets, AI system impact assessment scope, and AI project secure development in the AI policy or through cross-references to related policies.


**Control applies to:**
- Level 2
- Level 3




---

### AIGS-GA-08 - Regulatory Compliance

Demonstrate regulatory awareness by disclosing compliance status to the extent the open source AI project meets regulatory requirements and standards for the deployment of AI models and systems.



#### AIGS-GA-08-01

**Requirement:** To the extent the open source AI project meets regulatory requirements and standards for the deployment of AI models and systems, the project MUST disclose (non-)compliance.



**Control applies to:**
- Level 3




---

### AIGS-GA-09 - AI Bill of Materials

Enable downstream risk management by making structured disclosures of relevant model and data metadata available in machine-readable format while the project is active, to inform downstream risk-based controls.



#### AIGS-GA-09-01

**Requirement:** While active, structured disclosures of relevant model and data metadata relevant to AI project development and deployment MUST be made available in machine-readable format to inform downstream risk-based controls.



**Control applies to:**
- Level 3




---

## Supply Chain Security

Controls related to securing the AI supply chain, including data sourcing, component scanning, and model provenance.



### AIGS-SC-01 - Secure Data Sourcing

Protect data integrity by implementing controls to detect and prevent the unintentional sourcing of poisoned data, with documentation and disclosure of these controls.



#### AIGS-SC-01-01

**Requirement:** The project MUST implement controls to detect and prevent the unintentional sourcing of poisoned data. Controls MUST be documented and disclosed.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-SC-02 - Ethical Web Crawling

Ensure ethical data sourcing practices when web crawling is used, by respecting robots.txt requirements, clearly identifying the crawler, and disclosing web crawling activity in project documentation.



#### AIGS-SC-02-01

**Requirement:** To the extent web crawling is used to source data, web crawling activity MUST respect robots.txt requirements, the crawler MUST be clearly identified, and web crawling activity MUST be disclosed in project documentation.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-SC-03 - Change Control

Establish robust change management through use of a version control system and a documented process for reviewing and approving changes.



#### AIGS-SC-03-01

**Requirement:** The project MUST use a version control system and have a documented process for reviewing and approving changes.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-SC-04 - AI Components Scanning

Secure the supply chain by scanning embedded AI components as part of development workflows, including CI/CD pipelines, to prevent application security risks from malicious or otherwise compromised components.



#### AIGS-SC-04-01

**Requirement:** Embedded AI components MUST be scanned as part of development workflows, including CI/CD pipelines, to prevent application security risks from malicious or otherwise compromised components.



**Control applies to:**
- Level 2
- Level 3




---

### AIGS-SC-05 - Training Pipeline Integrity

Ensure training pipeline integrity by documenting and securing the training pipeline used to produce model artifacts, including dependencies, training code, and environment configuration.



#### AIGS-SC-05-01

**Requirement:** The project MUST document and secure the training pipeline used to produce model artifacts, including dependencies, training code, and environment configuration.



**Control applies to:**
- Level 2
- Level 3




---

### AIGS-SC-06 - Model Provenance

Establish model traceability by maintaining provenance records linking model artifacts, training datasets, training code, and training parameters.



#### AIGS-SC-06-01

**Requirement:** The project MUST maintain provenance records linking model artifacts, training datasets, training code, and training parameters.



**Control applies to:**
- Level 2
- Level 3




---

### AIGS-SC-07 - Reproducible Builds

Enable verification through support for reproducible model builds, where feasible, such that independent parties can recreate the model artifacts from disclosed inputs.



#### AIGS-SC-07-01

**Requirement:** Where feasible, the project SHOULD support reproducible model builds such that independent parties can recreate the model artifacts from disclosed inputs.


**Recommendation:** Provide sufficient documentation and tooling for independent parties to reproduce model builds from disclosed inputs.


**Control applies to:**
- Level 2
- Level 3




---

### AIGS-SC-08 - Model Integrity

Ensure artifact integrity by signing AI models, model files, and/or other release artifacts with cryptographic tools and verifying them before deployment to prevent the introduction of untrusted components.



#### AIGS-SC-08-01

**Requirement:** The project MUST sign AI models, model files, and/or other release artifacts with cryptographic tools and verify them before deployment to prevent the introduction of untrusted components.



**Control applies to:**
- Level 3




---

## Data Integrity

Controls related to data governance, security, access control, and leakage prevention.



### AIGS-DI-01 - Enforceable Data Policy

Maintain data provenance by tracking the origin and lineage of all data used for training and testing AI models.



#### AIGS-DI-01-01

**Requirement:** Track the origin and lineage of all data used for training and testing AI models.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-DI-02 - Data Security

Protect data through implementation of robust security controls to ensure confidentiality, integrity, and availability.



#### AIGS-DI-02-01

**Requirement:** The project MUST implement robust security controls to protect the confidentiality, integrity, and availability of data.



**Control applies to:**
- Level 1
- Level 2
- Level 3


#### AIGS-DI-02-02

**Requirement:** Minimum data security requirements MUST include authentication and authorization.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-DI-03 - Data Least Privilege Access

Implement secure access management through clearly documented data access controls that follow the principles of least privilege.



#### AIGS-DI-03-01

**Requirement:** Data access controls MUST be clearly documented and follow the principles of least privilege access.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-DI-04 - Sensitive and Copyright Data Leakage Prevention

Prevent data leakage by implementing controls and testing to avoid inadvertently revealing sensitive information (e.g., PII, proprietary data) from the training set.



#### AIGS-DI-04-01

**Requirement:** The project MUST implement controls and testing to prevent the project from inadvertently revealing sensitive information (e.g., PII, proprietary data) from its training set.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-DI-05 - Dataset Cataloguing

Maintain dataset traceability by tracking and versioning datasets used for training and evaluation to ensure reproducibility.



#### AIGS-DI-05-01

**Requirement:** Datasets used for training and evaluation MUST be tracked and versioned to ensure traceability and reproducibility.



**Control applies to:**
- Level 2
- Level 3




---

### AIGS-DI-06 - Dataset Risk Assessment

Mitigate dataset risks by assessing datasets for, at a minimum, legal and copyright risk, licensing compatibility with project intended use, and security risks (e.g., data poisoning, PII/SPI).



#### AIGS-DI-06-01

**Requirement:** The project MUST assess datasets for, at a minimum, legal and copyright risk, licensing compatibility with project intended use, and security risks (e.g., data poisoning, PII/SPI).



**Control applies to:**
- Level 2
- Level 3




---

### AIGS-DI-07 - Data Risk Minimization

Minimize data exposure by implementing controls to limit the use of unnecessary or high-risk data.



#### AIGS-DI-07-01

**Requirement:** The project MUST implement controls to limit the use of unnecessary or high-risk data.



**Control applies to:**
- Level 2
- Level 3




---

### AIGS-DI-08 - Full Audit Logging

Maintain comprehensive audit trail by logging read or write access to data, with timestamp, user id, action performed (read/write), and dataset id and version(s) accessed.



#### AIGS-DI-08-01

**Requirement:** Read or write access to data MUST be logged, with timestamp, user id, action performed (read/write), and dataset id and version(s) accessed.



**Control applies to:**
- Level 3




---

## Model Robustness

Controls related to adversarial testing, evaluations, and resilience of AI models.



### AIGS-MR-01 - Model Robustness

Strengthen model resilience through red teaming to screen for vulnerabilities to adversarial attacks, including prompt injection attacks.



#### AIGS-MR-01-01

**Requirement:** The project MUST undertake red teaming to screen for vulnerabilities to adversarial attacks, including prompt injection attacks.



**Control applies to:**
- Level 1
- Level 2
- Level 3


#### AIGS-MR-01-02

**Requirement:** To the extent possible, the project owner SHOULD disclose red teaming methodology and sufficient detail to satisfy downstream users of the model&#39;s robustness to a diverse set of adversarial attacks under deployment-like conditions.


**Recommendation:** Publish red teaming methodology and results with sufficient detail to inform downstream users about adversarial robustness.


**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-MR-02 - Evaluations

Validate model performance and security prior to deployment by testing against publicly available or manually-created, task-specific benchmarks to evaluate, quantify, and validate the resilience of the AI project against common vulnerabilities.



#### AIGS-MR-02-01

**Requirement:** Prior to deployment, the project team MUST test against publicly available or manually-created, task-specific benchmarks in order to evaluate, quantify, and validate the resilience of the AI project against common vulnerabilities.



**Control applies to:**
- Level 1
- Level 2
- Level 3


#### AIGS-MR-02-02

**Requirement:** To the extent possible, the project owner SHOULD disclose evaluation results and sufficient detail to satisfy downstream users of the model&#39;s reliability under a diverse set of deployment-like conditions and tasks.


**Recommendation:** Publish evaluation results and methodology to inform downstream users about model reliability under deployment-like conditions.


**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-MR-03 - Continuous Testing

Maintain ongoing security assurance by continuously implementing red teaming and evaluations prior to pushing system updates or model re-training into production.



#### AIGS-MR-03-01

**Requirement:** The project MUST continuously implement red teaming and evaluations prior to pushing system updates or model re-training into production.



**Control applies to:**
- Level 2
- Level 3




---

### AIGS-MR-04 - Regression Testing

Preserve security posture by maintaining regression tests for safety and security behaviors to prevent regression during updates.



#### AIGS-MR-04-01

**Requirement:** The project MUST maintain regression tests for safety and security behaviors to prevent regression during updates.



**Control applies to:**
- Level 2
- Level 3




---

### AIGS-MR-05 - Evaluation Coverage

Provide testing transparency by documenting the scope and limitations of testing.



#### AIGS-MR-05-01

**Requirement:** The project MUST document the scope and limitations of testing.



**Control applies to:**
- Level 2
- Level 3




---

### AIGS-MR-06 - Third-Party Red Teaming

Validate security through independent testing for adversarial robustness from trusted third-parties.



#### AIGS-MR-06-01

**Requirement:** The project MUST conduct or commission independent testing for adversarial robustness from trusted third-parties.



**Control applies to:**
- Level 3




---

### AIGS-MR-07 - Robustness Benchmark Testing

Ensure ongoing robustness by evaluating models used in the AI project at runtime against an evolving suite of standardized robustness benchmarks.



#### AIGS-MR-07-01

**Requirement:** The project MUST evaluate models used in the AI project at runtime against an evolving suite of standardized robustness benchmarks.



**Control applies to:**
- Level 3




---

## Deployment

Controls related to secure deployment, monitoring, incident response, and runtime safeguards.



### AIGS-DE-01 - Guardrails

Mitigate deployment risks by implementing guardrails to address identified risks, with minimum coverage for prompt injection attacks (input).



#### AIGS-DE-01-01

**Requirement:** The project MUST implement guardrails to mitigate identified risks, and at a minimum, implement guardrails to mitigate against prompt injection attacks (input).



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-DE-02 - Vulnerability Disclosure Mechanism

Enable responsible disclosure by establishing a defined communication channel or mechanism through which vulnerabilities and other security-related information may be disclosed to project contributors and maintainers.



#### AIGS-DE-02-01

**Requirement:** The project MUST have a defined communication channel or mechanism through which vulnerabilities and other security-related information MAY be disclosed to the project contributors and maintainers.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-DE-03 - Insecure Output Handling

Prevent downstream vulnerabilities by validating and sanitizing model outputs.



#### AIGS-DE-03-01

**Requirement:** The project MUST validate and sanitize model outputs to prevent downstream vulnerabilities.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-DE-04 - Deployment Infrastructure

Secure deployment infrastructure by ensuring access to models, tools and tool calling functionality, and general deployment scaffolding or infrastructure follows least-privilege principles.



#### AIGS-DE-04-01

**Requirement:** Access to models, tools and tool calling functionality, and general deployment scaffolding or infrastructure MUST follow least-privilege principles.



**Control applies to:**
- Level 2
- Level 3




---

### AIGS-DE-05 - Logging and Monitoring

Enable threat detection by implementing logging and monitoring for deployment environments, including tool use, to detect misuse or anomalies.



#### AIGS-DE-05-01

**Requirement:** The project MUST implement logging and monitoring for deployment environments, including tool use, to detect misuse or anomalies.



**Control applies to:**
- Level 2
- Level 3




---

### AIGS-DE-06 - Incident Response Procedures

Establish incident response capability by defining procedures for responding to security incidents involving the AI project in a timely manner.



#### AIGS-DE-06-01

**Requirement:** The project MUST define procedures for responding to security incidents involving the AI project in a timely manner.



**Control applies to:**
- Level 2
- Level 3




---

### AIGS-DE-07 - Project Maintenance

Ensure sustainable maintenance by documenting processes for releasing model updates and security patches.



#### AIGS-DE-07-01

**Requirement:** The project MUST document processes for releasing model updates and security patches.



**Control applies to:**
- Level 2
- Level 3




---

### AIGS-DE-08 - Runtime Safeguards

Protect production systems by including runtime safeguards in deployment environments for detecting and mitigating malicious inputs and outputs and unexpected or unauthorized system-level behavior.



#### AIGS-DE-08-01

**Requirement:** Deployment environments MUST include runtime safeguards for detecting and mitigating malicious inputs and outputs and unexpected or unauthorized system-level behavior.



**Control applies to:**
- Level 3




---

### AIGS-DE-09 - Kill Switch

Enable emergency response by implementing a security mechanism designed to immediately halt, disable, or contain an AI project if it is believed to behave in dangerous or unpredictable ways causing harm.



#### AIGS-DE-09-01

**Requirement:** The project MUST implement a security mechanism designed to immediately halt, disable, or contain an AI project if it is believed to behave in dangerous or unpredictable ways causing harm.



**Control applies to:**
- Level 3




---

## Transparency and Explainability

Controls related to system documentation, explainability, risk documentation, and incident reporting.



### AIGS-TE-01 - Explainability

Enhance understanding by providing explanations for design decisions that affect project performance and security, where feasible.



#### AIGS-TE-01-01

**Requirement:** Where feasible, the project MUST provide explanations for design decisions that affect project performance and security.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-TE-02 - System Documentation

Ensure system transparency through clear and comprehensive documentation of the AI system, including its intended use, limitations, and potential risks.



#### AIGS-TE-02-01

**Requirement:** Project contributors MUST maintain clear and comprehensive documentation for the AI system, including its intended use, limitations, and potential risks.



**Control applies to:**
- Level 1
- Level 2
- Level 3




---

### AIGS-TE-03 - Residual Risk Documentation

Promote informed decision-making by documenting residual risks and known limitations of safeguards.



#### AIGS-TE-03-01

**Requirement:** The project MUST document residual risks and known limitations of safeguards.



**Control applies to:**
- Level 3




---

### AIGS-TE-04 - Security Incident Reporting

Foster transparency and learning by maintaining a public record of safety and security incidents and ex post mitigations.



#### AIGS-TE-04-01

**Requirement:** The project MUST maintain a public record of safety and security incidents and ex post mitigations.



**Control applies to:**
- Level 3




---






## Acknowledgments

This document was developed, under the leadership of Derek Leist, thanks to contributions from technical experts across IBM Research, in addition to feedback and contributions from external collaborators including:
- [AIGS Baseline contributors](https://github.com/ibm/ai-security-baseline/graphs/contributors)

