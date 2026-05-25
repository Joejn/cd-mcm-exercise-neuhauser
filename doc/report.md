# Exercise Security K8s

**Name:** Jonas Neuhauser

**URL zum Repository:** https://github.com/Joejn/cd-mcm-exercise-neuhauser

## Aufgabe 1

### Trivy Ergebnis

**Vor den Verbesserungen**
![Trivy Scan](./assets/task_01_01.png)

Im Scan ist zu sehen, dass es 3 kleine, 5 mittlere und 2 große Sicherheitslücken im Image gibt. Dabei beeinflusst das Base-Image für den Deploy `alpine:3.19` den Sicherheitsbereich am meisten. Durch die Verwendung eines anderen Base-Images können diese Sicherheitslücken geschlossen werden.

**Mit dem Image Scratch als Base-Image**
![Trivy Scan](./assets/task_01_02.png)

## Aufgabe 2

### govulncheck Ergebnis

**Vor den Verbesserungen**
![Aufgabe 2 vor den Verbesserungen](./assets/task_02_01.png)

**Verbesserungen:** Die Go-Version in `go.mod` wurde auf `1.26.3` gesetzt.

**Nach den Verbesserungen**
![Aufgabe 2 nach den Verbesserungen](./assets/task_02_02.png)

## Aufgabe 3

### Deployment

**Logs**
![Logs](./assets/task_03_01.png)

**Healthcheck**
![Healthcheck](./assets/task_03_02.png)

### API Calls

**Alle Produkte abfragen**

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/products" -Method GET
```

![Alle Produkte abfragen](./assets/task_03_03.png)

**Ein Produkt abfragen**

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/products/1" -Method GET
```

![Ein Produkt abfragen](./assets/task_03_04.png)

**Produkt erstellen**

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/products" `
  -Method POST `
  -ContentType "application/json" `
  -Body '{"name":"Computer","price":999.99}'
```

![Produkt erstellen](./assets/task_03_05.png)

**Produkt aktualisieren**

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/products/1" `
  -Method PUT `
  -ContentType "application/json" `
  -Body '{"name":"Laptop","price":1299.99}'
```

![Produkt aktualisieren](./assets/task_03_06.png)

**Produkt löschen**

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/products/1" -Method DELETE
```

![Produkt löschen](./assets/task_03_07.png)

## Aufgabe 4

Lösung in K8S.md
