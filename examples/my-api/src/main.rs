//! Minimal API: GET /my-api returns JSON (status, uptime, board).
//! For Orange Pi Zero LTS / FreeBSD 15. Build: cargo build --release
//! Run: ./target/release/my-api (listens on 0.0.0.0:8080)

use std::io::{Read, Write};
use std::net::TcpListener;
use std::time::Instant;

const PORT: u16 = 8080;

fn main() -> std::io::Result<()> {
    let start = Instant::now();
    let addr = format!("0.0.0.0:{}", PORT);
    let listener = TcpListener::bind(&addr)?;
    eprintln!("my-api listening on http://{}", addr);

    for stream in listener.incoming().filter_map(Result::ok) {
        let _ = handle(stream, start);
    }
    Ok(())
}

fn handle(mut stream: std::net::TcpStream, start: Instant) -> std::io::Result<()> {
    let mut buf = [0u8; 512];
    let n = stream.read(&mut buf)?;
    let req = String::from_utf8_lossy(&buf[..n]);
    let first_line = req.lines().next().unwrap_or("");

    let (status, body) = if first_line.contains("GET /my-api") || first_line.contains("GET /my-api ") {
        let uptime_sec = start.elapsed().as_secs();
        let body = format!(
            r#"{{"status":"ok","uptime_sec":{},"board":"Orange Pi Zero LTS","endpoint":"/my-api"}}"#,
            uptime_sec
        );
        ("200 OK", body)
    } else {
        (r#"404 Not Found"#, r#"{"status":"not_found"}"#.to_string())
    };

    let response = format!(
        "HTTP/1.0 {}\r\nContent-Type: application/json\r\nContent-Length: {}\r\nConnection: close\r\n\r\n{}",
        status,
        body.len(),
        body
    );
    stream.write_all(response.as_bytes())?;
    stream.flush()?;
    Ok(())
}
