# Aufgabe 4

## Skalieren

![Skalieren](./assets/task_04_01.png)

## Health Checks

**Readiness vs Liveness Probe: what’s the difference?**
Unter Liveness versteht man, dass ein Pod gestartet ist. Das bedeutet aber noch nicht, dass er eingehende Anfragen bearbeiten kann. Zum Beispiel braucht die Anwendung im Pod eine Verbindung zur Datenbank, und erst wenn diese vorhanden ist, gilt der Pod als ready.

**What happens when each probe fails?**
Es gibt verschiedene Arten von Probes:

* **Liveness**

  * Wird verwendet, um festzustellen, wann ein Pod neu gestartet werden soll
* **Readiness**

  * Gibt an, ob der Container Anfragen verarbeiten kann
  * Wenn der Readiness-Probe einen Fehlerstatus zurückgibt, wird die Adresse des Pods vom EndpointSlice-Controller entfernt, wodurch keine Anfragen mehr an den Pod gesendet werden

**Why different initialDelaySeconds values?**
Der Wert `initialDelaySeconds` gibt an, wie lange nach dem Start ein Probe geprüft wird. Das `initialDelaySeconds` für den Liveness Probe ist häufig größer als für die Readiness Probe, weil es länger dauern kann, bis ein Pod komplett initialisiert ist.

## Ressourcen

**What happens if memory/CPU limit is exceeded?**
Wenn das Memory-Limit erreicht wird, wird der Pod gestoppt, weil durch Linux ein "Out of Memory" ausgelöst wird. Beim Erreichen des CPU-Limits wird der Pod nicht gestoppt, sondern der Container erhält lediglich weniger Rechenleistung.

**Why specify both requests and limits?**
Requests bestimmen, wie viel ein Pod **mindestens** benötigt. Das ist notwendig, damit Kubernetes entscheiden kann, auf welchem Node ein Pod gestartet werden kann. Die Limits geben an, wie viel ein Pod maximal beanspruchen darf.