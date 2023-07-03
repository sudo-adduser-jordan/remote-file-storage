import { cookies } from 'next/dist/client/components/headers';

const CREATE_PATH = 'http://localhost:8080/register';
const LOGIN_PATH = 'http://localhost:8080/login';
const LOG_OUT = 'http://localhost:8080/logout';

export async function createAccount(e: any, router: any): Promise<Boolean> {
  e.preventDefault();
  const form = e.target;
  const formData = new FormData(form);
  const headersList = {
    Accept: '*/*',
  };
  const response = await fetch(CREATE_PATH, {
    method: 'POST',
    headers: headersList,
    body: formData,
    credentials: 'include',
  });
  if (response.ok) {
    router.push('/protected/home');
    return false;
  } else {
    // Return error
    return true;
  }
}

export async function login(e: any, router: any): Promise<Boolean> {
  e.preventDefault();
  const form = e.target;
  const formData = new FormData(form);
  const headersList = {
    Accept: '*/*',
  };
  const response = await fetch(LOGIN_PATH, {
    method: 'POST',
    headers: headersList,
    body: formData,
    credentials: 'include',
  });
  if (response.ok) {
    router.push('/protected/home');
    return false;
  } else {
    // Return error
    return true;
  }
}

export function logout() {
  cookies().set({
    name: 'session_token',
    value: '',
    maxAge: 0,
    path: '/', // For all paths
    // expires: new Date('2016-10-05'),
  });
}

// Fix
export async function upload(e: any, router: any): Promise<Boolean> {
  e.preventDefault();
  const form = e.target;
  const formData = new FormData(form);
  const headersList = {
    Accept: '*/*',
  };
  const response = await fetch(LOGIN_PATH, {
    method: 'POST',
    headers: headersList,
    body: formData,
    credentials: 'include',
  });
  if (response.ok) {
    router.push('/protected/confirm');
    return false;
  } else {
    // Return error
    return true;
  }
}

// Fix
export async function download(e: any, router: any): Promise<Boolean> {
  e.preventDefault();
  const form = e.target;
  const formData = new FormData(form);
  const headersList = {
    Accept: '*/*',
  };
  const response = await fetch(LOGIN_PATH, {
    method: 'GET',
    headers: headersList,
    body: formData,
    credentials: 'include',
  });
  if (response.ok) {
    router.push('/protected/confirm');
    return false;
  } else {
    // Return error
    return true;
  }
}
