{ pkgs ? import <nixpkgs> {} }:

with pkgs; pkgs.mkShell rec {
  buildInputs = with pkgs; [
    alsa-lib
    xorg.libX11
    libGL
  ];
  LD_LIBRARY_PATH = pkgs.lib.makeLibraryPath buildInputs;
}
