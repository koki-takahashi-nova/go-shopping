'use client'

import { useState, useEffect } from 'react'
import { Cart, fetchCart, removeFromCart } from '@/lib/api'

export default function CartPage() {
  const [cart, setCart] = useState<Cart>({ items: [], total: 0 })
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    fetchCart().then(data => {
      setCart(data)
      setLoading(false)
    })
  }, [])

  const handleRemove = async (productId: number) => {
    const updated = await removeFromCart(productId)
    setCart(updated)
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
      <h1>ショッピングカート</h1>

      {cart.items.length === 0 ? (
        <div className="alert alert-info">カートは空です。</div>
      ) : (
        <>
          <table className="table">
            <thead>
              <tr>
                <th>商品名</th>
                <th>価格</th>
                <th>数量</th>
                <th>小計</th>
                <th></th>
              </tr>
            </thead>
            <tbody>
              {cart.items.map(item => (
                <tr key={item.product.id}>
                  <td>{item.product.name}</td>
                  <td>{item.product.price.toLocaleString()}円</td>
                  <td>{item.quantity}</td>
                  <td>{item.subtotal.toLocaleString()}円</td>
                  <td>
                    <button
                      className="btn btn-danger btn-sm"
                      onClick={() => handleRemove(item.product.id)}
                    >
                      削除
                    </button>
                  </td>
                </tr>
              ))}
            </tbody>
            <tfoot>
              <tr>
                <td colSpan={3} className="text-end">
                  <strong>合計:</strong>
                </td>
                <td colSpan={2}>{cart.total.toLocaleString()}円</td>
              </tr>
            </tfoot>
          </table>

          <div className="text-end mt-3">
            <a href="/" className="btn btn-secondary me-2">買い物を続ける</a>
            <a href="/checkout" className="btn btn-primary">注文する</a>
          </div>
        </>
      )}
    </div>
  )
}
