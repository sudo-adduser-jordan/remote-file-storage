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
    return true;
  }
}

// TODO: Delete Cookie
export async function logout(e: any, router: any) {}

export async function upload(e: any) {}

export async function download(e: any) {}
