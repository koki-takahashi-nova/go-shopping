import type { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'ショッピングサイト',
  description: 'Go + Next.js ショッピングアプリ',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="ja">
      <head>
        <link
          href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css"
          rel="stylesheet"
        />
      </head>
      <body>
        <nav className="navbar navbar-expand-lg navbar-light bg-light">
          <div className="container">
            <a className="navbar-brand" href="/">ショッピングサイト</a>
            <div className="navbar-nav ms-auto">
              <a className="nav-link" href="/cart">カート</a>
            </div>
          </div>
        </nav>
        <main>{children}</main>
        <script
          src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/js/bootstrap.bundle.min.js"
          async
        />
      </body>
    </html>
  )
}
