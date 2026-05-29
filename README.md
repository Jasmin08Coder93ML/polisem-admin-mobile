# Polisem Intelligent ERP Ekosistemasy

**Polisem** — islendik görnüşdäki retail (azyk, egin-eşik, elektronika, dermanhana, gurluşyk we ş.m.) we lomaý söwda dükanlaryny dolandyrmak üçin niýetlenen, Süni Intellekt goldawly, çeýe we döwrebap ERP (Enterprise Resource Planning) ulgamydyr.

Bu platforma **Akhasap** ýaly adaty programmalardan diňe bir maglumat saklamak däl, eýsem dynamic grafiki hasabatlar, salgyt (nalog) modulynyň awtomatizasiýasy we akylly AI Copilot integration arkaly düýbünden tapawutlanýar.

---

## 🛠 Tehnologiýalar (Tech Stack)

Ekosistema iň ýokary tizlik, howpsuzlyk we platformara durnuklylyk üçin iň döwrebap guraçlar bilen gurnaldy:

*   **Frontend (Mobil Goşundy):** Flutter (Dart) — Feature-First gurluşy we çeýe dizaýn ulgamy bilen.
*   **Backend (Arkat Engine):** Go (Golang) — Clean Architecture we ýokary öndürijilikli API süzgüçleri bilen.
*   **Maglumat Gory (Database):** PostgreSQL — Multi-shop we çeýe haryt atributlaryny saklamak üçin.
*   **Süni Intellekt Core:** Gemini API — Ses, tekst we wideo analizleri işlemek üçin.

---

## ✨ Esasy Aýratynlyklar (Key Features)

### 🏪 1. Ähli Söwda Ugurlary Üçin Çeýelik
Ulgam diňe bir gurluşyk dükanlary üçin däl, eýsem islendik söwda ugry üçin niýetlenendir. Ulanyjy dükanyň görnüşini saýlanda, haryt kartçyklary we atributlar (ölçeg, reňk, agram, möhlet) dükanyň ugruna görä dinamiki üýtgeýär.

### ✍️ 2. Elde Çalt Haryt Girizmek Ulgamy (Wizard UI)
Dükan eýeleriniň işini ýeňilleşdirmek üçin haryt goşmak interfeýsi minimalistik dizaýnda taýýarlandy. Ýekeje düwme bilen awtomatiki ştrihkod generirlenýär we kamera arkaly harydyň suraty çalt goşulýar.

### 📊 3. Wizual Hasabatlar we Nalog Moduly
*   Hasabatlar diňe bir adaty Excel görnüşinde çykman, eýsem **Line, Bar we Pie diagrammalar** (shemalar) arkaly dükanyň maliýe ýagdaýyny wizual görkezýär.
*   Döwlet salgyt kanunçylygyna laýyklykda awtomatiki **Nalog (Salgyt) hasabat şablonlaryny** taýýarlaýar.

### 🧠 4. AI Copilot (Akylly Süni Intellekt Kömekçisi)
*   **Sesli (Voice):** Ulanyjynyň "Şu günki jemi peýda näçe?" ýaly sesli soraglaryny tanap, bazadan maglumaty çekip berýär (Speech-to-Text).
*   **Tekstli (Text):** Dükan eýesine satuw meýilleri we bäsdeşlik barada maslahat berýär.
*   **Wideoly/Suratly (Vision):** Kamera arkaly harytlary, fakturalary ýa-da ştrihkodlary tanap, awtomatiki ulgama işleýär.

---

## 📁 Taslamanyň Gurluşy (Project Structure)

```text
polisem-platform/
├── backend/                  # GO BACKEND (Clean Architecture)
│   ├── cmd/api/main.go       # Başlangyç nokat
│   ├── internal/             # Auth, Shop, Products, Sales, Reports, AI Copilot
│   └── pkg/                  # Database drivers & Middlewares
└── frontend/                 # FLUTTER FRONTEND (Feature-Driven UI)
    ├── assets/               # Medýa we ikonlar
    └── lib/                  # Core & Features (Product management, POS, AI Assistant, Reports)
