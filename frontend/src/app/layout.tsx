export default function Layout({ children }) {
  return (
    <html lang="en">
      <head></head>
      <body className='container'>
        <main>{children}</main>
      </body>
    </html>
  );
}
