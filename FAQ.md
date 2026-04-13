# Open Source AI Project Governance and Security Baseline FAQ


## What's the purpose of the OSAIPGS Baseline?

The Open Source AI Project Governance and Security Baseline (OSAIPGS Baseline) is designed to act as a minimum set of requirements for AI projects relative to its maturity level. It builds on the [Open Source Project Security](https://baseline.openssf.org/versions/2025-10-10-checklist.md) (OSPS) Baseline to bring common security standards to open source AI projects, and recognizing the importance of governance controls in secure AI supply chains, incorporates new standards on data and model governance.

Overall, the Baseline aims to help maintainers, contributors, and other downstream users quickly understand and adopt fundamental governance and security steps. This includes establishing access-controlled environments for data acquisition and storage, setting up responsible disclosure policies, and maintaining base project hygiene. By meeting this Baseline, an open source AI project signals that it has taken the minimum measures required to reduce risk of common vulnerabilities and improve the overall trustworthiness of the project to its adopters and contributors.

## What problem is the OSAIPGS Baseline attempting to solve?

The open source AI community currently lacks easily adoptable standards, or minimum enforceable governance and security requirements, to standardize the way AI projects are securely built today. Because of this, it is often difficult to evaluate the risks of ingesting open source AI projects into other AI-native applications, software, or other downstream enterprise projects. In addition, open source AI projects often lack clear disclosures around governance and security practice, which further obfuscates risks from downstream users. This risks creating cascading vulnerabilities and hinders the application of truly risk-based controls at runtime. Emerging standards in this space are generally either too high-level to be practical, or else require bespoke tooling

## How is the OSAIPGS Baseline different from other open source security initiatives? How does the Baseline relate to established frameworks such as NIST AI RMF and ISO/IEC 42001? Is it a substitute for formal regulatory compliance or certification processes (ie. IEEE has a certification ‘badge')?

The OSAIPGS Baseline is a minimal governance and security checklist, based on project maturity levels. Enhanced measures can and should be taken where the risk warrants additional governance controls and security measures, particularly when deploying frontier models in mission critical systems or deployment environments which may warrant additional scrutiny to avoid harm to humans.

There may be some overlap between the OSAIPGS Baseline and other open source security initiatives and standards. To the extent possible, we have sought to align the Baseline to other emerging standards ([ISO 42001](https://www.iso.org/standard/42001)) and risk management frameworks (NIST RMF); however, these frameworks and standards are constantly evolving. Accordingly, project contributors and maintainers should regularly monitor the Baseline along with other relevant guidance, standards, risk management frameworks, and best practices for AI governance and security. Contributors to this Baseline included, but were not limited to AI governance experts, AI security professionals, AI alignment practitioners, and data management professionals.

## What should the Baseline not be used for?

The Baseline is not intended to be used as a one-size-fits-all scoring or grading mechanism for open source AI projects. Given the complexity of AI project, the diversity of downstream tasks or objectives, and the wide range of functionalities present in AI systems, it would not be appropriate to use the Baseline as a comparative tool between projects. Each project team that adopts the Baseline or attempts to meet the Baseline checklist is doing so, in good faith, to improve the governance and security posture of their open source AI project.

Feedback from the community is welcomed on ways to further improve the Baseline.

## As a maintainer of a project why should I bother with a Baseline?

The Baseline is meant as a foundation or starting point on which to implement good governance hygiene and security practice when structuring an AI project, from ideation to development to deployment. Achieving the criteria for an identified project maturity level signals to contributors and maintainers that the project is under development or deployment in a responsible, ethical way. Additionally, the Baseline can serve as a strong market signal to differentiate the project from other, unsecure or less secure alternatives, thereby incentivizing adoption and eliciting more contributions.

## What OSAIPGS Baseline level am I required to meet?

You are not required to meet any Baseline controls, unless you have a sponsoring organization (e.g. an open source foundation) that imposes a requirement. However, all projects are encouraged to adhere to Baseline Level 1 at a minimum because it establishes a "universal governance and security floor" for all open source AI projects. If you are a foundation that has some level or maturity criteria, we recommend you evaluate your lowest criteria tier for governance and security and adjust to align with the Baseline, where reasonable and appropriate.

Each maturity level maps to overall project maturity, as well as the risk profile of the intended downstream task and architecture of the system, application and/or tooling being built:

- Level 1: for any generative AI project with any number of maintainers or users and low risk profile
- Level 2: for any generative AI project that has at least 2 maintainers and a small number of consistent users or medium risk profile
- Level 3: for any generative AI project that has a large number of consistent users or large risk profile

Risk is a function of likelihood (probability) and consequence (impact/severity). Risk profiles should be assessed by the project team and documented, as required.

## How can I prove my project complies with the OSAIPGS Baseline?

Projects can self-attest OSAIPGS Baseline compliance. Eventually, tooling may be made available.

## How can I verify an upstream project's compliance with the OSAIPGS Baseline?

Many of the OSAIPGS Baseline controls are publicly observable. However, some of the controls are concerned with privileged or non-public development. Since the Baseline is designed for the developers, maintainers and contributors of an open source AI project, not the downstream users, if you need to verify the governance and security posture of an upstream project, users should accept self-attestation or else make an alternate arrangement with the project maintainers that meets your specific needs.

## Does OSAIPGS Baseline compliance expire?

The OSAIPGS Baseline compliance is a point-in-time status. We encourage projects using the OSAIPGS Baseline to include something to the effect of: "As of March 03, 2026, this project complies with OSAIPGS Baseline version 2025-03-06 Level 3" on the GitHub repository or project page.

## What is in scope for OSAIPGS Baseline?

The OSAIPGS Baseline seeks to address good governance and security hygiene - in particular, practices which govern and secure the ways of working, acquiring data, shipping models and systems, and deploying AI projects in a safe, controlled way.

To use an analogy, the Baseline is similar to construction site guidelines that require a construction companies to have practices in place that ensure safe construction of residential and commercial buildings: wearing hard hats, regularly logging construction site access, monitoring housing code requirements, auditing construction work with city residential and commercial building code, etc.

## How can I get involved in the OSAIPGS Baseline project?

The OSAIPGS Baseline project welcomes contributions in this Github Repository. For discussion, feel free to open an issue!
