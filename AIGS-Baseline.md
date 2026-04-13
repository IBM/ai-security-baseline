# Open Source AI Project Governance and Security Baseline

The Open Source AI Project Governance and Security Baseline (OSAIPGS Baseline) is designed to act as a minimum set of requirements for AI projects relative to its maturity level. It extends the principles of the [OpenSSF Security Baseline](https://baseline.openssf.org/) to address the unique challenges of developing, deploying, and managing Artificial Intelligence (AI) systems. It is designed to be a foundational guide for ensuring AI systems are secure, robust, transparent, and aligned with governance objectives.

For more information on the motive and purpose, see the [FAQ](FAQ.md).

## Versions

Current version v. 2026-03-06 (checklist)

## Guiding Principles

The OSAIPGS Baseline controls help AI project maintainers to understand governance and security best practices and expectations. Assessing a project's compliance against the controls helps maintainers and downstream users understand where the project meets minimum expected requirements for governance and security controls, and where improvements may be made. This gap assessment can inform risk-based controls at runtime of AI projects, and how consuming AI projects - including AI models, agents, and applications - into software products and other systems impact their own security and compliance objectives. Therefore, the OSAIPGS Baseline is:

- Focused: Controls contain MUST entries
- Realistic: Controls are practical for project maintainers to implement at the appropriate level for their project.
- Actionable: Controls provide specific recommendations.
- Meaningful: Controls have an impact on an AI project's security posture.

## Open Source AI Project Governance and Security Baseline Checklist

### Level 1

#### 1. Governance and Accountability

1.1 *Contributor Roles and Responsibilities*: The project MUST document the roles and responsibilities for maintainers and contributors to the AI project's lifecycle from development to deployment and ongoing monitoring.

1.2 *Model Governance:* The project MUST provide written, documented disclosure of the AI models ingested in, called, or otherwise used in the development or deployment of the open source AI project.

1.2a Minimum disclosure requirements MUST include the model name, developer, release date, license, model description, deployment status (as necessary), and URL to the open source model repository (if available) or model release announcement.

1.3 *Data Governance*: The project MUST disclose datasets ingested in or otherwise used in the development or deployment of the open source AI project. This includes datasets used for training, testing, and validation.

1.3a Minimum disclosure requirements for each dataset MUST include the name, country of origin, license, description, data processing (if any), and public URL (if available).

1.4 *Licensing*: While active, the project MUST specify a license for the model weights and model files.

1.5 *Ethics and Intended Use*: The project MUST define and document the intended use of the open source AI project, including the ethical principles that guide open source AI project development and deployment.

1.6 *Project* *Inventory*: The project MUST maintain an inventory of its components, tooling, systems and applications, disclosed within the README of the repository where the project is made available.

#### 2. Supply Chain Security

2.1 *Secure Data Sourcing*: The project MUST implement controls to detect and prevent the unintentional sourcing of poisoned data. Controls MUST be documented and disclosed.

2.2 *Ethical Web Crawling*: To the extent web crawling is used to source data, web crawling activity MUST respect robots.txt requirements, the crawler MUST be clearly identified, and web crawling activity MUST be disclosed in project documentation.

2.3 *Change Control*: The project MUST use a version control system and have a documented process for reviewing and approving changes.

#### 3. Data Integrity

3.1 *Enforceable* *Data Policy*: Track the origin and lineage of all data used for training and testing AI models.

3.2 *Data Security*: The project MUST implement robust security controls to protect the confidentiality, integrity, and availability of data.

3.2a Minimum data security requirements MUST include: authentication and authorization.

3.3 *Data Least Privilege Access*: Data access controls MUST be clearly documented and follow the principles of least privilege access.

3.4 *Sensitive and Copyright Data Leakage Prevention*: The project MUST implement controls and testing to prevent the project from inadvertently revealing sensitive information (e.g., PII, proprietary data) from its training set.

#### 4. Model Robustness

4.1 *Model Robustness*: The project MUST undertake red teaming to screen for vulnerabilities to adversarial attacks, including prompt injection attacks.

4.1a *Model Robustness* *Disclosures:* To the extent possible, the project owner SHOULD disclose red teaming methodology and sufficient detail to satisfy downstream users of the model's robustness to a diverse set of adversarial attacks under deployment-like conditions.

4.2 *Evaluations:* Prior to deployment, the project team MUST test against publicly available or manually-created, task-specific benchmarks in order to evaluate, quantify, and validate the resilience of the AI project against common vulnerabilities.

4.2a *Evaluations Disclosures:* To the extent possible, the project owner SHOULD disclose evaluation results and sufficient detail to satisfy downstream users of the model's reliability under a diverse set of deployment-like conditions and tasks.

#### 5. Deployment

5.1 *Guardrails*: The project MUST implement guardrails to mitigate identified risks, and at a minimum, implement guardrails to mitigate against prompt injection attacks (input).

5.2 *Vulnerability Disclosure Mechanism:* The project MUST have a defined communication channel or mechanism through which vulnerabilities and other security-related information may be disclosed to the project contributors and maintainers.

5.3 *Insecure Output Handling*: The project MUST validate and sanitize model outputs to prevent downstream vulnerabilities.

#### 6. Transparency and Explainability

6.1 *Explainability*: Where feasible, the project MUST provide explanations for design decisions that affect project performance and security.

6.2 *System Documentation*: Project contributors MUST maintain clear and comprehensive documentation for the AI system, including its intended use, limitations, and potential risks.

### Level 2

#### 1. Governance and Accountability

1.7 *AI Policy*: The project MUST define an AI policy with i) principles that guide all activities of the organization related to AI and ii) processes for handling deviations and exceptions to policy.

1.7a *AI Policy Requirements:* The AI policy should consider topic-specific aspects where necessary to provide additional guidance or provide cross-references to other security and governance policies dealing with these core aspects: i) AI resources and assets; ii) AI system impact assessment scope; and iii) AI project secure development.

#### 2. Supply Chain Security

2.4 *AI Components Scanning:* Embedded AI components MUST be scanned as part of development workflows, including CI/CD pipelines, to prevent application security risks from malicious or otherwise compromised components.

2.5 *Training Pipeline Integrity:* The project MUST document and secure the training pipeline used to produce model artifacts, including dependencies, training code, and environment configuration.

2.6 *Model Provenance:* The project MUST maintain provenance records linking: model artifacts, training datasets, training code, and training parameters.

2.7 *Reproducible Builds:* Where feasible, the project SHOULD support reproducible model builds such that independent parties can recreate the model artifacts from disclosed inputs.

#### 3. Data Integrity

3.5 *Dataset Cataloguing*: Datasets used for training and evaluation MUST be tracked and versioned to ensure traceability and reproducibility.

3.6 *Dataset Risk Assessment*: The project MUST assess datasets for, at a minimum: legal and copyright risk, licensing compatibility with project intended use, and security risks (e.g., data poisoning, PII/SPI).

3.7 *Data Risk Minimization*: The project MUST implement controls to limit the use of unnecessary or high-risk data.

#### 4. Model Robustness

4.3 *Continuous Testing*: The project MUST continuously implement red teaming and evaluations prior to pushing system updates or model re-training into production.

4.4 *Regression Testing:* The project MUST maintain regression tests for safety and security behaviors to prevent regression during updates.

4.5 *Evaluation Coverage:* The project MUST document the scope and limitations of testing.

#### 5. Deployment

5.4 *Deployment Infrastructure*: Access to models, tools and tool calling functionality, and general deployment scaffolding or infrastructure MUST follow least-privilege principles.

5.5 *Logging and Monitoring:* The project MUST implement logging and monitoring for deployment environments, including tool use, to detect misuse or anomalies.

5.6 *Incident Response Procedures:* The project MUST define procedures for responding to security incidents involving the AI project in a timely manner.

5.7 *Project Maintenance:* The project MUST document processes for releasing model updates and security patches.

#### 6. Transparency and Explainability

TBD

### Level 3

#### 1. Governance and Accountability

1.8 *Regulatory Compliance*: To the extent the open source AI project meets regulatory requirements and standards for the deployment of AI models and systems, the project MUST disclose (non-)compliance.

1.9 *AI Bill of Materials:* While active, structured disclosures of relevant model and data metadata relevant to AI project development and deployment MUST be made available in machine-readable format to inform downstream risk-based controls.

#### 2. Supply Chain Security

2.8 *Model Integrity*: The project MUST sign AI models, model files, and/or other release artifacts with cryptographic tools and verify them before deployment to prevent the introduction of untrusted components.

#### 3. Data Integrity

3.6 *Full Audit Logging:* Read or write access to data must be logged, with timestamp, user id, action performed (read/write), and dataset id and version(s) accessed.

#### 4. Model Robustness

4.6 *Third-Party Red Teaming:* The project MUST conduct or commission independent testing for adversarial robustness from trusted third-parties.

4.7 *Robustness Benchmark Testing:* The project MUST evaluate models used in the AI project at runtime against an evolving suite of standardized robustness benchmarks.

#### 5. Deployment

5.8 *Runtime Safeguards:* Deployment environments MUST include runtime safeguards for detecting and mitigating malicious inputs and outputs and unexpected or unauthorized system-level behavior.

5.9 *Kill Switch*: The project MUST implement a security mechanism designed to immediately halt, disable, or contain an AI project if it is believed to behave in dangerous or unpredictable ways causing harm.

#### 6. Transparency and Explainability

6.3 *Residual* *Risk Documentation:* The project MUST document residual risks and known limitations of safeguards.

6.4 *Security Incident Reporting:* The project MUST maintain a public record of safety and security incidents and ex post mitigations.


## Copyright

IBM Corporation
