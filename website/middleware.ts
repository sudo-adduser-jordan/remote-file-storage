import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';

export function middleware(request: NextRequest) {
  console.log(request.nextUrl.pathname);
  const session = request.cookies.get('session_token')?.value;
  console.log(session);
  if (session == undefined) {
    return NextResponse.redirect(new URL('/', request.url));
  }
}

export const config = {
  matcher: '/protected/:path*',
};
