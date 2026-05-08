export interface Product {
  id: number
  name: string
  description: string
  price: number
  category: string
}

export interface CartItem {
  product: Product
  quantity: number
  subtotal: number
}

export interface Cart {
  items: CartItem[]
  total: number
}

export interface Customer {
  name: string
  email: string
  address: string
  phoneNumber: string
}

export interface Order {
  id: number
  customerId: number
  customer: Customer
  orderDate: string
  orderDetails: OrderDetail[]
  totalAmount: number
}

export interface OrderDetail {
  id: number
  orderId: number
  productId: number
  product: Product
  quantity: number
  price: number
}

const API_BASE = '/api'

async function apiFetch<T>(path: string, init?: RequestInit): Promise<T> {
  const res = await fetch(`${API_BASE}${path}`, {
    credentials: 'include',
    ...init,
    headers: {
      'Content-Type': 'application/json',
      ...init?.headers,
    },
  })
  if (!res.ok) {
    const body = await res.json().catch(() => ({}))
    throw new Error(body.error ?? `API error: ${res.status}`)
  }
  return res.json()
}

export const fetchProducts = (): Promise<Product[]> =>
  apiFetch('/products')

export interface SearchParams {
  keyword?: string
  minPrice?: number
  maxPrice?: number
}

export const searchProducts = (params: SearchParams): Promise<Product[]> => {
  const q = new URLSearchParams()
  if (params.keyword) q.set('keyword', params.keyword)
  if (params.minPrice != null) q.set('minPrice', String(params.minPrice))
  if (params.maxPrice != null) q.set('maxPrice', String(params.maxPrice))
  return apiFetch(`/products/search?${q}`)
}

export const fetchProductsByCategory = (category: string): Promise<Product[]> =>
  apiFetch(`/products/category/${encodeURIComponent(category)}`)

export const filterLow = (): Promise<Product[]> => apiFetch('/products/filter/low')
export const filterMid = (): Promise<Product[]> => apiFetch('/products/filter/mid')
export const filterHigh = (): Promise<Product[]> => apiFetch('/products/filter/high')

export const fetchCart = (): Promise<Cart> => apiFetch('/cart')

export const addToCart = (productId: number): Promise<Cart> =>
  apiFetch(`/cart/add/${productId}`, { method: 'POST' })

export const removeFromCart = (productId: number): Promise<Cart> =>
  apiFetch(`/cart/remove/${productId}`, { method: 'DELETE' })

export const placeOrder = (customer: Customer): Promise<{ order: Order }> =>
  apiFetch('/orders', {
    method: 'POST',
    body: JSON.stringify(customer),
  })
