# gRPC Practice 2026
Chapter 2: https://hun-bot2.github.io/study/gRPC/gRPC-basic

## 출처
**"gRPC: Up and Running"**  
저자: Kasun Indrasiri, Danesh Kuruppu  
출판사: O'Reilly Media (2020)  
ISBN: 978-1492058335

## 목표
- gRPC 기본 개념 및 통신 패턴 학습
- Go를 사용한 gRPC 서비스 구현 실습
- Production 환경에서의 gRPC 배포 및 운영 방법 습득
- 최신 버전(2026) 호환성 문제 해결 및 문서화

## 학습 기간
2026년 1월 6일 ~ 진행 중

---

## Book Chapters (Go Implementation)

### Chapter 01 - Introduction to gRPC [FIN]
No code samples required

### Chapter 02 - Getting Started with gRPC [FIN]
- ProductInfo Service (Basic gRPC)

### Chapter 03 - gRPC Communication Patterns
- Order Service (Unary, Server Streaming, Client Streaming, Bidirectional Streaming)

### Chapter 04 - gRPC Under the Hood
No code samples required

### Chapter 05 - gRPC Beyond the Basics
- Interceptors
- Deadline
- Cancellation
- Compression
- Keepalive
- Metadata
- Error Handling
- Load Balancing
- Multiplexing

### Chapter 06 - Secured gRPC
- Secure Channel
- Mutual TLS Channel
- Basic Authentication
- Token Based Authentication

### Chapter 07 - Running gRPC in Production
- Continuous Integration
- Deploy in Docker
- Deploy in Kubernetes
- OpenCensus Metrics
- OpenCensus Tracing
- OpenTracing
- Prometheus

### Chapter 08 - The gRPC Ecosystem
- gRPC Gateway
- Server Reflection

---

## Repository Structure

```folder
gRPC-Practice-2026/
├── productinfo/          # Chapter 02: Getting Started with gRPC
└── Readmd.md
```

## Reference
원본 샘플 코드: [grpc-up-and-running/samples](https://github.com/grpc-up-and-running/samples)