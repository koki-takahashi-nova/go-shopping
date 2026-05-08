'use client'

import { useState, useEffect } from 'react'
import { useRouter } from 'next/navigation'
import { Cart, Customer, fetchCart, placeOrder } from '@/lib/api'

export default function CheckoutPage() {
  const router = useRouter()
  const [cart, setCart] = useState<Cart>({ items: [], total: 0 })
  const [loading, setLoading] = useState(true)
  const [submitting, setSubmitting] = useState(false)
  const [error, setError] = useState('')
  const [form, setForm] = useState<Customer>({
    name: '',
    email: '',
    address: '',
    phoneNumber: '',
  })

  useEffect(() => {
    fetchCart().then(data => {
      setCart(data)
      setLoading(false)
    })
  }, [])

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
    setForm(prev => ({ ...prev, [e.target.name]: e.target.value }))
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setSubmitting(true)
    setError('')
    try {
      await placeOrder(form)
      router.push('/order/confirmation')
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : '注文に失敗しました')
      setSubmitting(false)
    }
  }

  if (loading) {
    return (
      <div className="container mt-4 text-center">
        <div className="spinner-border" role="status" />
      </div>
    )
  }

  return (
    <div className="container mt-4">
      <h1>注文情報入力</h1>

      {error && <div className="alert alert-danger">{error}</div>}

      <div className="row">
        <div className="col-md-6">
          <h3>お客様情報</h3>
          <form onSubmit={handleSubmit}>
            <div className="mb-3">
              <label className="form-label">お名前</label>
              <input
                type="text"
                className="form-control"
                name="name"
                value={form.name}
                onChange={handleChange}
                required
              />
            </div>
            <div className="mb-3">
              <label className="form-label">メールアドレス</label>
              <input
                type="email"
                className="form-control"
                name="email"
                value={form.email}
                onChange={handleChange}
                required
              />
            </div>
            <div className="mb-3">
              <label className="form-label">住所</label>
              <textarea
                className="form-control"
                name="address"
                value={form.address}
                onChange={handleChange}
                rows={3}
                required
              />
            </div>
            <div className="mb-3">
              <label className="form-label">電話番号</label>
              <input
                type="tel"
                className="form-control"
                name="phoneNumber"
                value={form.phoneNumber}
                onChange={handleChange}
                required
              />
            </div>
            <button
              type="submit"
              className="btn btn-primary"
              disabled={submitting}
            >
              {submitting ? '処理中...' : '注文を確定する'}
            </button>
          </form>
        </div>

        <div className="col-md-6">
          <h3>注文内容</h3>
          <table className="table">
            <thead>
              <tr>
                <th>商品名</th>
                <th>数量</th>
                <th>小計</th>
              </tr>
            </thead>
            <tbody>
              {cart.items.map(item => (
                <tr key={item.product.id}>
                  <td>{item.product.name}</td>
                  <td>{item.quantity}</td>
                  <td>{item.subtotal.toLocaleString()}円</td>
                </tr>
              ))}
            </tbody>
            <tfoot>
              <tr>
                <td colSpan={2} className="text-end">
                  <strong>合計:</strong>
                </td>
                <td>{cart.total.toLocaleString()}円</td>
              </tr>
            </tfoot>
          </table>
        </div>
      </div>
    </div>
  )
}
