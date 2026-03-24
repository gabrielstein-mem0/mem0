const https = require("https");
const fs = require("fs");
const path = require("path");
const { execFileSync } = require("child_process");

const VERSION = require("./package.json").version;

const PLATFORM_MAP = { darwin: "darwin", linux: "linux", win32: "windows" };
const ARCH_MAP = { arm64: "arm64", x64: "amd64" };

const platform = PLATFORM_MAP[process.platform];
const arch = ARCH_MAP[process.arch];

if (!platform || !arch) {
  console.error(`Unsupported platform: ${process.platform} ${process.arch}`);
  process.exit(1);
}

const isWindows = process.platform === "win32";
const ext = isWindows ? ".zip" : ".tar.gz";
const archiveName = `mem0_${VERSION}_${platform}_${arch}${ext}`;
const url = `https://github.com/mem0ai/mem0/releases/download/v${VERSION}/${archiveName}`;

const binDir = path.join(__dirname, "bin");
const archivePath = path.join(__dirname, archiveName);
const binaryName = isWindows ? "mem0-bin.exe" : "mem0-bin";
const binaryPath = path.join(binDir, binaryName);

fs.mkdirSync(binDir, { recursive: true });

function download(url) {
  return new Promise((resolve, reject) => {
    https.get(url, (res) => {
      if (res.statusCode >= 300 && res.statusCode < 400 && res.headers.location) {
        return download(res.headers.location).then(resolve, reject);
      }
      if (res.statusCode !== 200) {
        return reject(new Error(`Download failed: HTTP ${res.statusCode}\n  ${url}`));
      }
      const chunks = [];
      res.on("data", (chunk) => chunks.push(chunk));
      res.on("end", () => resolve(Buffer.concat(chunks)));
      res.on("error", reject);
    }).on("error", reject);
  });
}

async function main() {
  console.log(`Downloading mem0 v${VERSION} for ${platform}/${arch}...`);

  const data = await download(url);
  fs.writeFileSync(archivePath, data);

  const extractedName = isWindows ? "mem0.exe" : "mem0";
  const extractedPath = path.join(binDir, extractedName);

  try {
    if (isWindows) {
      execFileSync("powershell", [
        "-Command",
        `Expand-Archive -Force -Path '${archivePath}' -DestinationPath '${binDir}'`,
      ]);
    } else {
      execFileSync("tar", ["xzf", archivePath, "-C", binDir, extractedName]);
    }

    fs.renameSync(extractedPath, binaryPath);
    fs.chmodSync(binaryPath, 0o755);
    console.log(`Installed mem0 to ${binaryPath}`);
  } finally {
    fs.unlinkSync(archivePath);
  }
}

main().catch((err) => {
  console.error(`Failed to install mem0: ${err.message}`);
  process.exit(1);
});
