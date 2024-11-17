import ky from "ky"

const apiClient = ky.create({
  prefixUrl: "/api/v2",
  headers: {
    "Content-Type": "application/json",
  },
})

export default apiClient
