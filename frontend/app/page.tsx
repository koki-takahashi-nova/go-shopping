'use client'

import { useState, useEffect, useCallback } from 'react'
import {
  Product,
  fetchProducts,
  searchProducts,
  filterLow,
  filterMid,
  filterHigh,
  addToCart,
} from '@/lib/api'

export default function HomePage() {
  const [products, setProducts] = useState<Product[]>([])
  const [keyword, setKeyword] = useState('')
  const [minPrice, setMinPrice] = useState('')
  const [maxPrice, setMaxPrice] = useState('')
  const [resultCount, setResultCount] = useState<number | null>(null)
  const [loading, setLoading] = useState(false)
  const [cartMessage, setCartMessage] = useState('')

  const loadProducts = useCallback(async () => {
    setLoading(true)
    const data = await fetchProducts()
    setProducts(data)
    setResultCount(null)
    setLoading(false)
  }, [])

  useEffect(() => {
    loadProducts()
  }, [loadProducts])

  const handleSearch = async (e: React.FormEvent) => {
    e.preventDefault()
    setLoading(true)
    const data = await searchProducts({
      keyword: keyword || undefined,
      minPrice: minPrice ? Number(minPrice) : undefined,
      maxPrice: maxPrice ? Number(maxPrice) : undefined,
    })
    setProducts(data)
    setResultCount(data.length)
    setLoading(false)
  }

  const handleFilter = async (fn: () => Promise<Product[]>, label: string) => {
    setLoading(true)
    const data = await fn()
    setProducts(data)
    setResultCount(data.length)
    setLoading(false)
  }

  const handleReset = () => {
    setKeyword('')
    setMinPrice('')
    setMaxPrice('')
    loadProducts()
  }

  const handleAddToCart = async (id: number) => {
    await addToCart(id)
    setCartMessage('カートに追加しました')
    setTimeout(() => setCartMessage(''), 2000)
  }

  return (
    <div className="container mt-4">
      <h1>商品一覧</h1>

      {cartMessage && (
        <div className="alert alert-success alert-dismissible" role="alert">
          {cartMessage}
        </div>
      )}

      <form onSubmit={handleSearch} className="mb-4">
        <div className="row g-3 align-items-end">
          <div className="col-md-4">
            <label htmlFor="keyword" className="form-label">商品名キーワード</label>
            <input
              type="text"
              className="form-control"
              id="keyword"
              value={keyword}
              onChange={e => setKeyword(e.target.value)}
              placeholder="キーワードを入力"
            />
          </div>
          <div className="col-md-3">
            <label htmlFor="minPrice" className="form-label">最低価格（円）</label>
            <input
              type="number"
              className="form-control"
              id="minPrice"
              value={minPrice}
              onChange={e => setMinPrice(e.target.value)}
              placeholder="下限なし"
              min={0}
            />
          </div>
          <div className="col-md-3">
            <label htmlFor="maxPrice" className="form-label">最高価格（円）</label>
            <input
              type="number"
              className="form-control"
              id="maxPrice"
              value={maxPrice}
              onChange={e => setMaxPrice(e.target.value)}
              placeholder="上限なし"
              min={0}
            />
          </div>
          <div className="col-md-2">
            <button type="submit" className="btn btn-primary w-100">検索</button>
          </div>
        </div>
      </form>

      <div className="mb-4">
        <button
          className="btn btn-outline-secondary me-2"
          onClick={() => handleFilter(filterLow, '～10,000円')}
        >
          ～10,000円
        </button>
        <button
          className="btn btn-outline-secondary me-2"
          onClick={() => handleFilter(filterMid, '10,001円～50,000円')}
        >
          10,001円～50,000円
        </button>
        <button
          className="btn btn-outline-secondary me-2"
          onClick={() => handleFilter(filterHigh, '50,001円～')}
        >
          50,001円～
        </button>
        <button className="btn btn-outline-secondary" onClick={handleReset}>
          リセット
        </button>
      </div>

      {resultCount !== null && (
        <div className="mb-3">
          {resultCount === 0 ? (
            <p className="text-muted">条件に一致する商品が見つかりませんでした</p>
          ) : (
            <p className="text-muted">{resultCount}件見つかりました</p>
          )}
        </div>
      )}

      {loading ? (
        <div className="text-center py-5">
          <div className="spinner-border" role="status" />
        </div>
      ) : (
        <div className="row">
          {products.map(product => (
            <div key={product.id} className="col-md-4 mb-4">
              <div className="card h-100">
                <div className="card-body d-flex flex-column">
                  <h5 className="card-title">{product.name}</h5>
                  <p className="card-text flex-grow-1">{product.description}</p>
                  <p className="card-text">
                    <strong>価格: </strong>
                    {product.price.toLocaleString()}円
                  </p>
                  <p className="card-text">
                    <strong>カテゴリ: </strong>
                    {product.category}
                  </p>
                  <button
                    className="btn btn-primary mt-auto"
                    onClick={() => handleAddToCart(product.id)}
                  >
                    カートに追加
                  </button>
                </div>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}
