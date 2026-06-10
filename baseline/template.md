# Open Source AI Project Governance and Security Baseline

Version: {{ .Catalog.Metadata.Version }}

{: .warning}
Not for production use.

<!-- A button for returning to the top of the page -->
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
<!-- That's enough button stuff for now -->

* Contents
{:toc}

## Overview

The Open Source AI Project Governance and Security Baseline (AIGS Baseline) is designed to act as a minimum set of requirements for AI projects relative to its maturity level. It extends the principles of the [OpenSSF Security Baseline](https://baseline.openssf.org/) to address the unique challenges of developing, deploying, and managing Artificial Intelligence (AI) systems. It is designed to be a foundational guide for ensuring AI systems are secure, robust, transparent, and aligned with governance objectives.

For more information on the motive and purpose, see the [FAQ](FAQ.md).

For more information on the project and to make contributions, visit the [GitHub repo](https://github.com/ossf/security-baseline).

---

## Controls Overview
{{ range .Catalog.Metadata.ApplicabilityGroups }}
{{- $req := . }}
* [{{ $req.Title }}]({{ $req.Title | asLink }}): {{ $req.Description }}
{{- end }}


### Level 1
{{ range .Catalog.Controls }}
{{- range .AssessmentRequirements }}
{{- $req := . }}
{{- if and (not (isRetired .State)) (maxLevel .Applicability 1) }}
**[{{ $req.Id }}]({{ $req.Id | asLink }})**: {{ $req.Text | addLinks }}
{{ end }}
{{- end }}
{{- end }}

### Level 2
{{ range .Catalog.Controls }}
{{- range .AssessmentRequirements }}
{{- $req := . }}
{{- if and (not (isRetired .State)) (maxLevel .Applicability 2) }}
**[{{ $req.Id }}]({{ $req.Id | asLink }})**: {{ $req.Text | addLinks }}
{{ end }}
{{- end }}
{{- end }}

### Level 3
{{ range .Catalog.Controls }}
{{- range .AssessmentRequirements }}
{{- $req := . }}
{{- if and (not (isRetired .State)) (maxLevel .Applicability 3) }}
**[{{ $req.Id }}]({{ $req.Id | asLink }})**: {{ $req.Text | addLinks }}
{{ end }}
{{- end }}
{{- end }}

{{ range .Catalog.Groups }}

## {{ .Title }}

{{ .Description }}

{{ range controlsForGroup $.Catalog.Controls .Id }}

### {{ .Id }} - {{ .Title | addLinks | collapseNewlines }}

{{ .Objective }}

{{ range .AssessmentRequirements }}

#### {{ .Id }}

{{ if isRetired .State -}}
{{ .Text | collapseNewlines }}
{{ else -}}
**Requirement:** {{ .Text | addLinks | collapseNewlines }}

{{ if .Recommendation }}
**Recommendation:** {{ .Recommendation }}
{{ end }}

**Control applies to:**
{{ range .Applicability }}- {{ . | applicabilityTitle }}
{{ end }}

{{- end -}}

{{ end }}

{{ if  .Guidelines }}
#### External Framework Relations
  {{ range .Guidelines }}
  - **{{ .ReferenceId | addLinks }}**: {{ range $index, $entry := .Entries }}{{ if $index }}, {{ end }}{{ $entry.ReferenceId }}{{ end }}
  {{- end }}
{{ end }}

---

{{- end }}
{{- end }}

{{ if .Catalog.Metadata.MappingReferences -}}
## External Frameworks

Controls within this document may relate to the following external frameworks:

| ID | Title | Version | Description |
|----|-------|---------|-------------|
{{ range .Catalog.Metadata.MappingReferences -}}
| {{ .Id }} | [{{ .Title }}]({{ .Url }}) | {{ .Version }} | {{ .Description }} |
{{ end }}
---
{{ end }}

{{ if .Lexicon }}
## Lexicon
{{ range .Lexicon }}

### {{ .Term }}

{{ .Definition }}
{{ end}}

{{ if .References }}
**References:**
{{ range .References }}
  - {{.}}
{{ end -}}
{{ end -}}
---
{{ end }}


## Acknowledgments

This document was developed, under the leadership of Derek Leist, thanks to contributions from technical experts across IBM Research, in addition to feedback and contributions from external collaborators including:
- [AIGS Baseline contributors](https://github.com/ibm/ai-security-baseline/graphs/contributors)
{{ range .Lexicon }}
[{{ .Term }}]: {{ .Term | asLink }}
{{- end }}
