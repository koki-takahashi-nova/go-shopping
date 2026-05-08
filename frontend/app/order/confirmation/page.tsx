export default function OrderConfirmationPage() {
  return (
    <div className="container mt-4">
      <div className="alert alert-success">
        <h1 className="alert-heading">ご注文ありがとうございます！</h1>
        <p>ご注文を受け付けました。</p>
        <p>ご登録いただいたメールアドレスに確認メールをお送りしました。</p>
      </div>
      <div className="text-center mt-4">
        <a href="/" className="btn btn-primary">トップページに戻る</a>
      </div>
    </div>
  )
}
