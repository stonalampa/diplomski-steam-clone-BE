import http from 'k6/http';
import { check } from 'k6';

const URL = 'http://localhost:3030/api/login'; // Update with your server URL
const payload = JSON.stringify({
  email: 'user@test.com',
  password: 'test123',
});

export default function loginSimultaneously() {
  const parsedEnvVar = parseInt(__ENV.userCount);
  const userCount = isNaN(parsedEnvVar) ? 2 : parsedEnvVar;

  const requests = [];
  for (let i = 0; i < userCount; i++) {
    requests.push({
      method: 'POST',
      url: URL,
      headers: { 'Content-Type': 'application/json' },
      body: payload,
      tags: { name: `User ${i + 1}` },
    });
  }

  const responses = http.batch(requests);

  responses.forEach((res, i) => {
    check(res, {
      'Login Successful': (r) => r.status === 200,
    });

    if (res.status === 200) {
      console.log(`Login successful for User ${i + 1}`);
    } else {
      console.error(`Login failed for User ${i + 1}`);
    }
  });
}
