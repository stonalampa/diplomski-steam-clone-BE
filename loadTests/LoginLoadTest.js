import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
  const payload = JSON.stringify({
    email: 'user@test.com',
    password: 'test123',
  });

  const res = http.post('http://localhost:3030/api/login', payload, {
    headers: { 'Content-Type': 'application/json' },
  });

  if (res.status === 200) {
    console.log('Login successful');
  } else {
    console.error('Login failed');
  }

  sleep(1); // Sleep for 1 second between requests
}

export const options = {
  vus: 10,
  duration: '30s',
};