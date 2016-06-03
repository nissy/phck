Summary: PHCK is system health checker web server.
Name: phck
Version: 0.1
Release: 1
License: GNU General Public License, Version 3
Group: Applications/Services
URL: https://github.com/ngc224/phck

Source0: https://github.com/ngc224/phck/releases/download/%{version}/phck-%{version}_linux_amd64.tar.gz
Source1: pchk.conf
Source2: pchk.service

BuildRoot: %{_tmppath}/%{name}-root

%description
PHCK is system health checker web server.

%prep
%setup -q

%install
rm -rf $RPM_BUILD_ROOT

mkdir -p $RPM_BUILD_ROOT%{_sbindir}
install -m 755 phck $RPM_BUILD_ROOT%{_sbindir}

mkdir -p $RPM_BUILD_ROOT%{_sysconfdir}/pchk
install -m 644 %SOURCE1 $RPM_BUILD_ROOT%{_sysconfdir}/pchk/pchk.conf

mkdir -p $RPM_BUILD_ROOT%{_unitdir}
install -m 644 %SOURCE2 $RPM_BUILD_ROOT%{_unitdir}/pchk.service

%clean
rm -rf $RPM_BUILD_ROOT

%files
%defattr(-,root,root)
%{_sbindir}/pchk
%config(noreplace) %{_sysconfdir}/pchk/pchk.conf
%{_unitdir}/nginx.service

%changelog
* Mon Mar 17 2013 Yoshihiko Nishida <nishida@ngc224.org> - 7-2.el7.centos
- Build for CentOS-7.2