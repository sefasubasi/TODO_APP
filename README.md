# Erişim Linki:
https://todo-app-93kc.onrender.com

---

# ✅ TODO_APP - GoLang RESTful API

Bu proje, Go dilinde geliştirilmiş JWT tabanlı kullanıcı kimlik doğrulama ve rol bazlı erişim kontrolü içeren bir **To-Do yönetim sistemi API**'sidir. Kullanıcılar giriş yaptıktan sonra kendi görevlerini yönetebilir; admin rolündeki kullanıcılar ise tüm kullanıcı verilerini görüntüleyebilir.

---

## 📌 Özellikler

- 🔐 JWT (JSON Web Token) ile güvenli kimlik doğrulama
- 👤 Kullanıcı rolleri: `user` ve `admin`
- 🗃️ Kendi görevlerini listeleme, oluşturma, güncelleme, silme
- 🛡️ Admin yetkisiyle tüm kullanıcı görevlerini görüntüleme
- 🧠 Mock veri deposu ile çalışır (veritabanı gerekmez)
- 🏗️ MVC mimarisine uygun dizin yapısı

---

## 📁 Proje Yapısı

TODO_APP/
├── controller/ -> API uç noktaları (Login, ToDo)
├── middleware/ -> JWT ve Admin rol doğrulama
├── mock/ -> Sahte veri deposu (in-memory)
├── model/ -> Veri modelleri (User, Role, TodoList)
├── router/ -> Route tanımlamaları
├── service/ -> İş mantığı (Login vs.)
├── utils/ -> JWT oluşturma/doğrulama
├── main.go -> Ana dosya
├── go.mod -> Go modül tanımı
└── README.md -> Bu dosya

---

## 🔧 Kurulum

### Gereksinimler:
- [Go](https://go.dev/dl/) 1.18 veya üzeri

### Kurulum Adımları:

1. Proje dizinine git:
```bash
cd TODO_APP
go mod tidy
go run .


---

## Giriş Bilgileri (Mock Kullanıcılar)
| Rol   | Kullanıcı Adı | Şifre      |
| ----- | ------------- | ---------- |
| User  | `user1`       | `123456`   |
| Admin | `admin1`      | `admin123` |

---

## Kullanılabilir Endpoint’ler

### 🔑1.Kullanıcı Girişi (Login),
POST /login

Request Body (JSON):
{
  "username": "admin1",
  "password": "admin123"
}

Response:
{
  "token": "eyJhbGciOiJIUzI1NiIsInR..."
}

### 👤 2. Kullanıcı To-Do İşlemleri (Token Gereklidir)
Tüm bu isteklerde Authorization header'ı gereklidir:
Authorization: Bearer <token>
| Method | Endpoint    | Açıklama                             |
| ------ | ----------- | ------------------------------------ |
| GET    | `/todo/my`  | Kendi To-Do listelerini getirir      |
| POST   | `/todo`     | Yeni bir liste oluşturur             |
| PUT    | `/todo/:id` | Bir listeyi günceller                |
| DELETE | `/todo/:id` | Listeyi soft-delete olarak işaretler |

### 🛡️ 3. Admin Yetkili İşlemler
| Method | Endpoint       | Açıklama                                     |
| ------ | -------------- | -------------------------------------------- |
| GET    | `/admin/todos` | Tüm kullanıcıların to-do listelerini getirir |

---

## 🧪 Postman ile Test Etme
1-Giriş yap (POST /login)

2-Dönen token'ı kopyala

3- Diğer isteklerde Authorization başlığına şunu ekle:
Authorization: Bearer <token>
{
  "username": "admin1",
  "password": "admin123"
}


### 🔐 1. Login → Token al
POST http://localhost:8080/login

### 📋 2. Tüm kullanıcıların listelerini getir (admin)
GET http://localhost:8080/admin/todos
Authorization: Bearer <token>

---


📝 Lisans
MIT License © 2025
Bu proje açık kaynak olarak sunulmuştur. İstediğiniz gibi kullanabilir, geliştirebilir ve dağıtabilirsiniz.

📬 İletişim
sefasubasii@gmail.com
