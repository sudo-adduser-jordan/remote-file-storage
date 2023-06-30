const login_path = 'http://localhost:8080/login';

export async function login(e: any) {
  e.preventDefault();

  const form = e.target;
  const formData = new FormData(form);

  const headersList = {
    Accept: '*/*',
    Authorization: 'Basic dXNlcjE6cGFzcw==',
  };

  const response = await fetch(login_path, {
    method: 'POST',
    headers: headersList,
    body: formData,
  });
  console.log(response);
  console.log(response.ok);

  if (response.ok) {
    // router.push('/home');
  }
}

export async function logout(e: any) {}

export async function createAccount(e: any) {
  e.preventDefault();

  const form = e.target;
  const formData = new FormData(form);

  const headersList = {
    Accept: '*/*',
  };

  const response = await fetch(login_path, {
    method: 'POST',
    headers: headersList,
    body: formData,
  });
  console.log(response);
  console.log(response.ok);

  if (response.ok) {
    // router.push('/home');
  }
}
