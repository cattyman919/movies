import './NavigationBar.css';

export function NavigationBar() {
  return (
    <nav className="navbar">
      <div className="navbar-title">MovieDB</div>
      <div className="navbar-links">
        <button className="signin-button">Sign In</button>
      </div>
    </nav>
  );
}
