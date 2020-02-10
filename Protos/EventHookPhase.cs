// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: EventHookPhase.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace ULZAsset.ProtoMod {

  /// <summary>Holder for reflection information generated from EventHookPhase.proto</summary>
  public static partial class EventHookPhaseReflection {

    #region Descriptor
    /// <summary>File descriptor for EventHookPhase.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static EventHookPhaseReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "ChRFdmVudEhvb2tQaGFzZS5wcm90bxIIVUxaUHJvdG8qkQYKDkV2ZW50SG9v",
            "a1BoYXNlEhEKDWdhbWVzZXRfc3RhcnQQABIUChBzdGFydF90dXJuX3BoYXNl",
            "EAoSHAoYcmVmaWxsX2FjdGlvbl9jYXJkX3BoYXNlEBQSGwoXcmVmaWxsX2V2",
            "ZW50X2NhcmRfcGhhc2UQGRIYChRtb3ZlX2NhcmRfZHJvcF9waGFzZRAeEhgK",
            "FGRldGVybWluZV9tb3ZlX3BoYXNlECgSFQoRZmluaXNoX21vdmVfcGhhc2UQ",
            "MhIWChJjaGFyYV9jaGFuZ2VfcGhhc2UQPBIgChxkZXRlcm1pbmVfY2hhcmFf",
            "Y2hhbmdlX3BoYXNlEEYSHQoZYXR0YWNrX2NhcmRfZHJvcF9waGFzZV9wQRBQ",
            "Eh4KGmRlZmVuY2VfY2FyZF9kcm9wX3BoYXNlX3BBEFoSIwofZGV0ZXJtaW5l",
            "X2JhdHRsZV9wb2ludF9waGFzZV9wQRBkEhoKFmJhdHRsZV9yZXN1bHRfcGhh",
            "c2VfcEEQbhITCg9kYW1hZ2VfcGhhc2VfcEEQeBIeChpkZWFkX2NoYXJhX2No",
            "YW5nZV9waGFzZV9wQRB9EikKJGRldGVybWluZV9kZWFkX2NoYXJhX2NoYW5n",
            "ZV9waGFzZV9wQRCCARIeChlhdHRhY2tfY2FyZF9kcm9wX3BoYXNlX3BCEIwB",
            "Eh8KGmRlZmVuY2VfY2FyZF9kcm9wX3BoYXNlX3BCEJYBEiQKH2RldGVybWlu",
            "ZV9iYXR0bGVfcG9pbnRfcGhhc2VfcEIQoAESGwoWYmF0dGxlX3Jlc3VsdF9w",
            "aGFzZV9wQhCqARIUCg9kYW1hZ2VfcGhhc2VfcEIQtAESKQokZGV0ZXJtaW5l",
            "X2RlYWRfY2hhcmFfY2hhbmdlX3BoYXNlX3BCEL4BEh8KGmRlYWRfY2hhcmFf",
            "Y2hhbmdlX3BoYXNlX3BCEMgBEiYKIWRldGVybWluZV9kZWFkX2NoYXJhX2No",
            "YW5nZV9waGFzZRDSARIWChFmaW5pc2hfdHVybl9waGFzZRDcARIQCgtnYW1l",
            "c2V0X2VuZBDmASo+Cg1FdmVudEhvb2tUeXBlEgsKB0luc3RhbnQQABIKCgZC",
            "ZWZvcmUQARIJCgVBZnRlchACEgkKBVByb3h5EANCFKoCEVVMWkFzc2V0LlBy",
            "b3RvTW9kYgZwcm90bzM="));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { },
          new pbr::GeneratedClrTypeInfo(new[] {typeof(global::ULZAsset.ProtoMod.EventHookPhase), typeof(global::ULZAsset.ProtoMod.EventHookType), }, null, null));
    }
    #endregion

  }
  #region Enums
  public enum EventHookPhase {
    /// <summary>
    //// start gameset
    /// </summary>
    [pbr::OriginalName("gameset_start")] GamesetStart = 0,
    /// <summary>
    //// -- turn lifecycle
    /// </summary>
    [pbr::OriginalName("start_turn_phase")] StartTurnPhase = 10,
    /// <summary>
    //// ---- Draw phase (ddp)
    /// </summary>
    [pbr::OriginalName("refill_action_card_phase")] RefillActionCardPhase = 20,
    [pbr::OriginalName("refill_event_card_phase")] RefillEventCardPhase = 25,
    /// <summary>
    //// ---- Move phase
    /// </summary>
    [pbr::OriginalName("move_card_drop_phase")] MoveCardDropPhase = 30,
    [pbr::OriginalName("determine_move_phase")] DetermineMovePhase = 40,
    [pbr::OriginalName("finish_move_phase")] FinishMovePhase = 50,
    [pbr::OriginalName("chara_change_phase")] CharaChangePhase = 60,
    [pbr::OriginalName("determine_chara_change_phase")] DetermineCharaChangePhase = 70,
    /// <summary>
    //// ---- Atk/Def Phase A
    /// </summary>
    [pbr::OriginalName("attack_card_drop_phase_pA")] AttackCardDropPhasePA = 80,
    [pbr::OriginalName("defence_card_drop_phase_pA")] DefenceCardDropPhasePA = 90,
    [pbr::OriginalName("determine_battle_point_phase_pA")] DetermineBattlePointPhasePA = 100,
    /// <summary>
    ///roll dice
    /// </summary>
    [pbr::OriginalName("battle_result_phase_pA")] BattleResultPhasePA = 110,
    [pbr::OriginalName("damage_phase_pA")] DamagePhasePA = 120,
    [pbr::OriginalName("dead_chara_change_phase_pA")] DeadCharaChangePhasePA = 125,
    [pbr::OriginalName("determine_dead_chara_change_phase_pA")] DetermineDeadCharaChangePhasePA = 130,
    /// <summary>
    //// ---- Atk/Def Phase B
    /// </summary>
    [pbr::OriginalName("attack_card_drop_phase_pB")] AttackCardDropPhasePB = 140,
    [pbr::OriginalName("defence_card_drop_phase_pB")] DefenceCardDropPhasePB = 150,
    [pbr::OriginalName("determine_battle_point_phase_pB")] DetermineBattlePointPhasePB = 160,
    /// <summary>
    ///roll dice
    /// </summary>
    [pbr::OriginalName("battle_result_phase_pB")] BattleResultPhasePB = 170,
    [pbr::OriginalName("damage_phase_pB")] DamagePhasePB = 180,
    [pbr::OriginalName("determine_dead_chara_change_phase_pB")] DetermineDeadCharaChangePhasePB = 190,
    /// <summary>
    //// ------ change char when the charecter dead
    /// </summary>
    [pbr::OriginalName("dead_chara_change_phase_pB")] DeadCharaChangePhasePB = 200,
    [pbr::OriginalName("determine_dead_chara_change_phase")] DetermineDeadCharaChangePhase = 210,
    /// <summary>
    //// -- endof turn lifecycle
    /// </summary>
    [pbr::OriginalName("finish_turn_phase")] FinishTurnPhase = 220,
    /// <summary>
    //// endof game set
    /// </summary>
    [pbr::OriginalName("gameset_end")] GamesetEnd = 230,
  }

  public enum EventHookType {
    [pbr::OriginalName("Instant")] Instant = 0,
    [pbr::OriginalName("Before")] Before = 1,
    [pbr::OriginalName("After")] After = 2,
    [pbr::OriginalName("Proxy")] Proxy = 3,
  }

  #endregion

}

#endregion Designer generated code
