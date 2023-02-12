import {els, req} from "./dom";

export function initPermissions(t: HTMLSelectElement, s: HTMLSelectElement) {
  permissionsTeamToggle(t.value !== "");
  permissionsSprintToggle(s.value !== "");
}

export function permissionsTeamToggle(set: boolean) {
  req(".permission-config-team").style.display = set ? "block" : "none";
}

export function permissionsSprintToggle(set: boolean) {
  req(".permission-config-sprint").style.display = set ? "block" : "none";
}

export function loadPermsForm(frm: HTMLFormElement) {
  type StrMap = { [key: string]: string };
  const ret: StrMap = {};
  for (const el of els<HTMLInputElement>(".perm-option", frm)) {
    console.log(el);
    if (el.checked) {
      ret[el.name] = el.value;
    }
  }
  return ret;
}
