export const server = {
    baseURL: 'http://127.0.0.1:8181',
    addStatUrl: 'http://localhost:5050'
}

export function authHeader() {
    // return authorization header with jwt token
    let user = JSON.parse(localStorage.getItem('user'));

    if (user && user.token) {
        return { 'Authorization': 'Bearer ' + user.token };
    } else {
        return {};
    }
}
export const getters = {
    isAuthenticated: state => !!state.token,
    authStatus: state => state.status,
}