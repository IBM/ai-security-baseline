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


**[AIGS-GA-02-01](#aigs-ga-02-01)**: Minimum disclosure requirements MUST include the model name, developer, release date, license, model description, deployment status (as necessary), and URL to the open source model repository (if available) or model release announcement.

**[AIGS-GA-03-01](#aigs-ga-03-01)**: Minimum disclosure requirements for each dataset MUST include the name, country of origin, license, description, data processing (if any), and public URL (if available).

**[AIGS-GA-04-01](#aigs-ga-04-01)**: While active, the project MUST specify a license for the model weights and model files.

**[AIGS-GA-05-01](#aigs-ga-05-01)**: The project MUST define and document the intended use of the open source AI project, including the ethical principles that guide open source AI project development and deployment.

**[AIGS-GA-06-01](#aigs-ga-06-01)**: The project MUST maintain an inventory of its components, tooling, systems and applications, disclosed within the README of the repository where the project is made available.

**[AIGS-SC-01-01](#aigs-sc-01-01)**: The project MUST implement controls to detect and prevent the unintentional sourcing of poisoned data. Controls MUST be documented and disclosed.

**[AIGS-SC-02-01](#aigs-sc-02-01)**: To the extent web crawling is used to source data, web crawling activity MUST respect robots.txt requirements, the crawler MUST be clearly identified, and web crawling activity MUST be disclosed in project documentation.

**[AIGS-SC-03-01](#aigs-sc-03-01)**: The project MUST use a version control system and have a documented process for reviewing and approving changes.

**[AIGS-DI-01-01](#aigs-di-01-01)**: Track the origin and lineage of all data used for training and testing AI models.

**[AIGS-DI-02-01](#aigs-di-02-01)**: Minimum data security requirements MUST include authentication and authorization.

**[AIGS-DI-03-01](#aigs-di-03-01)**: Data access controls MUST be clearly documented and follow the principles of least privilege access.

**[AIGS-DI-04-01](#aigs-di-04-01)**: The project MUST implement controls and testing to prevent the project from inadvertently revealing sensitive information (e.g., PII, proprietary data) from its training set.

**[AIGS-MR-01-01](#aigs-mr-01-01)**: To the extent possible, the project owner SHOULD disclose red teaming methodology and sufficient detail to satisfy downstream users of the model&#39;s robustness to a diverse set of adversarial attacks under deployment-like conditions.

**[AIGS-MR-02-01](#aigs-mr-02-01)**: To the extent possible, the project owner SHOULD disclose evaluation results and sufficient detail to satisfy downstream users of the model&#39;s reliability under a diverse set of deployment-like conditions and tasks.

**[AIGS-DE-01-01](#aigs-de-01-01)**: The project MUST implement guardrails to mitigate identified risks, and at a minimum, implement guardrails to mitigate against prompt injection attacks (input).

**[AIGS-DE-02-01](#aigs-de-02-01)**: The project MUST have a defined communication channel or mechanism through which vulnerabilities and other security-related information MAY be disclosed to the project contributors and maintainers.

**[AIGS-DE-03-01](#aigs-de-03-01)**: The project MUST validate and sanitize model outputs to prevent downstream vulnerabilities.

**[AIGS-TE-01-01](#aigs-te-01-01)**: Where feasible, the project MUST provide explanations for design decisions that affect project performance and security.

**[AIGS-TE-02-01](#aigs-te-02-01)**: Project contributors MUST maintain clear and comprehensive documentation for the AI system, including its intended use, limitations, and potential risks.


### Level 2

**[AIGS-GA-07-01](#aigs-ga-07-01)**: The AI policy SHOULD consider topic-specific aspects where necessary to provide additional guidance or provide cross-references to other security and governance policies dealing with these core aspects: AI resources and assets, AI system impact assessment scope, and AI project secure development.

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









## Acknowledgments

This document was developed, under the leadership of Derek Leist, thanks to contributions from technical experts across IBM Research, in addition to feedback and contributions from external collaborators including:
- [AIGS Baseline contributors](https://github.com/ibm/ai-security-baseline/graphs/contributors)

