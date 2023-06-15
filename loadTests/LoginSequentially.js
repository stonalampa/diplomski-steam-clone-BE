import http from 'k6/http';
import { check } from 'k6';

const URL = 'http://localhost:3030/api/login'; // Update with your server URL
const payload = JSON.stringify({
  email: 'user@test.com',
  password: 'test123',
});

export default function loginSequentially() {
  const parsedEnvVar = parseInt(__ENV.userCount);
  const userCount = isNaN(parsedEnvVar) ? 2 : parsedEnvVar;

  for (let i = 0; i < userCount; i++) {
    const res = http.post(URL, payload, {
      headers: { 'Content-Type': 'application/json' },
    });

    check(res, {
      'Login Successful': (r) => r.status === 200,
    });

    if (res.status === 200) {
      console.log(`Login successful for User ${i + 1}`);
    } else {
      console.error(`Login failed for User ${i + 1}`);
    }
  }
}
